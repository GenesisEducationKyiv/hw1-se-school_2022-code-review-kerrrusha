package storage

import (
	"os"
)

type FileAccessable struct {
	Path string
}

func (f *FileAccessable) AccessFileRead() *os.File {
	file, err := os.Open(f.Path)
	if err != nil {
		panic(err)
	}
	return file
}

func (f *FileAccessable) AccessFileWrite() *os.File {
	file, err := os.Create(f.Path)
	if err != nil {
		panic(err)
	}
	return file
}
