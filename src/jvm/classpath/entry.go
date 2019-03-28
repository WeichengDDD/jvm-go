package classpath

import (
	"os"
	"strings"
)

const pathListSeparator = string(os.PathListSeparator)

type Entry interface {
	readClass(className string) (bytes []byte, entry Entry, err error)
	String() string
}

func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
	}

	return newDirEntry(path)
}
