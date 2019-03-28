package classpath

import (
	"io/ioutil"
	"path/filepath"
)

type DirEntry struct {
	absDir string
}

func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{absDir}
}

func (self *DirEntry) String() string {
	return self.absDir
}

func (self *DirEntry) readClass(className string) (bytes []byte, entry Entry, err error) {
	fileName := filepath.Join(self.absDir)
	data, err := ioutil.ReadFile(fileName)
	return data, self, err
}
