package comparator

import (
	"Users/dbogatyreva/GolandProjects/1C-task/internal/my_math"
	"Users/dbogatyreva/GolandProjects/1C-task/internal/utils"
	"sync"
	"sync/atomic"
)

func CompareDirs(directory1 string, directory2 string, threshold float64) (*Result, error) {
	filesFromFirstDir, err := utils.GetAllFiles(directory1)
	if err != nil {
		return nil, err
	}

	filesFromSecondDir, err := utils.GetAllFiles(directory2)
	if err != nil {
		return nil, err
	}

	result := MakeResult()

	filesFromSecondDirFound := make([]atomic.Bool, len(filesFromSecondDir), len(filesFromSecondDir))

	var wg sync.WaitGroup
	for _, fileFromFirstDir := range filesFromFirstDir {
		wg.Add(1)
		f := fileFromFirstDir
		go compareFiles(f, filesFromSecondDir, filesFromSecondDirFound, &wg, result, threshold)
	}
	wg.Wait()

	for i, file := range filesFromSecondDir {
		if !filesFromSecondDirFound[i].Load() {
			result.addNotFoundFileFromSecondDir(file.Name)
		}
	}

	return result, nil
}

func compareFiles(fileFromFirstDir utils.File, filesFromSecondDir []utils.File, filesFromSecondDirFound []atomic.Bool,
	wg *sync.WaitGroup, result *Result, threshold float64) {
	defer wg.Done()

	fileFromFirstDirFound := false

	for j, fileFromSecondDir := range filesFromSecondDir {
		if fileFromFirstDir.Hash == fileFromSecondDir.Hash {
			result.addIdenticalFile(fileFromFirstDir.Name, fileFromSecondDir.Name)
			fileFromFirstDirFound = true
			filesFromSecondDirFound[j].Store(true)
			continue
		}

		largerSize := len(fileFromFirstDir.Content)
		if len(fileFromSecondDir.Content) > largerSize {
			largerSize = len(fileFromSecondDir.Content)
		}

		thresholdDistance := int(float64(len(fileFromFirstDir.Content)) * (1 - threshold))
		distance := levenshteinDistance(fileFromFirstDir.Content, fileFromSecondDir.Content)

		if distance <= thresholdDistance {
			result.addSimilarFile(fileFromFirstDir.Name, fileFromSecondDir.Name)
			fileFromFirstDirFound = true
			filesFromSecondDirFound[j].Store(true)
		}
	}

	if !fileFromFirstDirFound {
		result.addNotFoundFileFromFirstDir(fileFromFirstDir.Name)
	}
}

// levenshteinDistance возвращает расстояние Левенштейна между двумя строками
func levenshteinDistance(s1, s2 []byte) int {
	m, n := len(s1), len(s2)

	dist := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dist[i] = make([]int, n+1)
		dist[i][0] = i
	}

	for j := 0; j <= n; j++ {
		dist[0][j] = j
	}

	for j := 1; j <= n; j++ {
		for i := 1; i <= m; i++ {
			if s1[i-1] == s2[j-1] {
				dist[i][j] = dist[i-1][j-1]
			} else {
				dist[i][j] = my_math.Min(dist[i-1][j]+1, dist[i][j-1]+1, dist[i-1][j-1]+1)
			}
		}
	}

	return dist[m][n]
}
