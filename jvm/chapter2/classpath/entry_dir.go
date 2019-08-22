package classpath

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

type DirEntry struct {
	absDir string //用于存放目录的绝对路径
}

func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	fmt.Printf("--------DirEntry absDir : %s -------\n", absDir)
	return &DirEntry{absDir}
}

func (dirEntry *DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(dirEntry.absDir, className)
	fmt.Printf("--------DirEntry fileName : %s -------\n", fileName)
	data, err := ioutil.ReadFile(fileName)
	return data, dirEntry, err
}

func (dirEntry *DirEntry) String() string {
	return dirEntry.absDir
}
