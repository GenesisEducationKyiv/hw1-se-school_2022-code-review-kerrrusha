package fileStorage

import (
	"os"

	"github.com/kerrrusha/btc-api/api/dataAccess/storage"
)

type fileReader struct {
	*storage.FileAccessable
}

func (reader *fileReader) Read() []byte {
	file := reader.AccessFileRead()

	databyte, err := os.ReadFile(reader.Path)
	if err != nil {
		panic(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	return databyte
}

func CreateFileReader(filepath string) *fileReader {
	return &fileReader{
		FileAccessable: &storage.FileAccessable{Path: filepath},
	}
}
