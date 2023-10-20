package comparator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompareDirs_IdenticalFiles(t *testing.T) {
	directory1 := "test_files/dir1"
	directory2 := "test_files/dir2"
	threshold := 0.33

	result, err := CompareDirs(directory1, directory2, threshold)

	if err != nil {
		t.Errorf("Error comparing directories: %s", err.Error())
	}

	identicalFiles := result.GetIdenticalFiles()
	assert.Equal(t, 1, len(identicalFiles))

	notFoundFilesFromFirstDir := result.GetNotFoundFilesFromFirstDir()
	notFoundFilesFromSecondDir := result.GetNotFoundFilesFromSecondDir()

	assert.Equal(t, 2, len(notFoundFilesFromFirstDir))
	assert.Equal(t, 1, len(notFoundFilesFromSecondDir))

	similarFiles := result.GetSimilarFiles()

	assert.Equal(t, 1, len(similarFiles))
}
