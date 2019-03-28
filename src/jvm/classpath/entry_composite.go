package classpath

import (
	"errors"
	"strings"
)

type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry {
	compsiteEntry := []Entry{}

	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(path)
		compsiteEntry = append(compsiteEntry, entry)
	}

	return compsiteEntry
}

func (self CompositeEntry) readClass(className string) (bytes []byte, entry Entry, err error) {
	for _, entry := range self {
		data, from, err := entry.readClass(className)
		if err == nil {
			return data, from, nil
		}
	}

	return nil, nil, errors.New("class not found: " + className)
}

func (self CompositeEntry) String() string {
	strs := make([]string, len(self))
	for i, entry := range self {
		strs[i] = entry.String()
	}

	return strings.Join(strs, pathListSeparator)
}
