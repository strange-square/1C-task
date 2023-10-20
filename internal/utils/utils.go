package utils

import (
	"crypto/md5"
	"os"
	"path/filepath"
)

type File struct {
	Name    string
	Content []byte
	Hash    [16]byte
}

func GetAllFiles(directory string) ([]File, error) {
	var files []File

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			content, err := getFileContent(path)
			if err != nil {
				return err
			}
			files = append(files, File{
				Name:    path,
				Content: content,
				Hash:    md5.Sum(content),
			})
		}
		return nil
	})

	return files, err
}

func getFileContent(file string) ([]byte, error) {
	content, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	return content, nil
}
