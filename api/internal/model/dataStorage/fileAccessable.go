package dataStorage

import (
	"os"

	"github.com/kerrrusha/BTC-API/api/internal/errorUtils"
)

type FileAccessable struct {
	Path string
}

func (f *FileAccessable) AccessFileRead() *os.File {
	file, err := os.Open(f.Path)
	errorUtils.CheckForError(err)
	return file
}

func (f *FileAccessable) AccessFileWrite() *os.File {
	file, err := os.Create(f.Path)
	errorUtils.CheckForError(err)
	return file
}
