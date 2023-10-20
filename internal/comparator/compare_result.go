package comparator

import (
	"sync"
)

type Result struct {
	identicalFiles             map[string][]string
	similarFiles               map[string][]string
	notFoundFilesFromFirstDir  []string
	notFoundFilesFromSecondDir []string
	mutex                      sync.RWMutex
}

func MakeResult() *Result {
	return &Result{
		identicalFiles:             make(map[string][]string, 0),
		similarFiles:               make(map[string][]string, 0),
		notFoundFilesFromFirstDir:  make([]string, 0),
		notFoundFilesFromSecondDir: make([]string, 0),
	}
}

func (c *Result) GetIdenticalFiles() map[string][]string {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.identicalFiles
}

func (c *Result) GetSimilarFiles() map[string][]string {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.similarFiles
}

func (c *Result) GetNotFoundFilesFromFirstDir() []string {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.notFoundFilesFromFirstDir
}

func (c *Result) GetNotFoundFilesFromSecondDir() []string {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.notFoundFilesFromSecondDir
}

func (c *Result) addIdenticalFile(fileName1 string, fileName2 string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.identicalFiles[fileName1] = append(c.identicalFiles[fileName1], fileName2)
}

func (c *Result) addSimilarFile(fileName1 string, fileName2 string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.similarFiles[fileName1] = append(c.identicalFiles[fileName1], fileName2)
}

func (c *Result) addNotFoundFileFromFirstDir(fileName string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.notFoundFilesFromFirstDir = append(c.notFoundFilesFromFirstDir, fileName)
}

func (c *Result) addNotFoundFileFromSecondDir(fileName string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.notFoundFilesFromSecondDir = append(c.notFoundFilesFromSecondDir, fileName)
}
