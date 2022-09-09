package fileStorage

import (
	"io/ioutil"

	"github.com/kerrrusha/BTC-API/api/internal/errorUtils"
	"github.com/kerrrusha/BTC-API/api/internal/model/dataStorage"
)

type fileReader struct {
	*dataStorage.FileAccessable
}

func (reader *fileReader) Read() []byte {
	file := reader.AccessFileRead()

	databyte, err := ioutil.ReadFile(reader.Path)
	errorUtils.CheckForError(err)

	defer file.Close()

	return databyte
}

func CreateFileReader(filepath string) *fileReader {
	return &fileReader{
		FileAccessable: &dataStorage.FileAccessable{Path: GetProjPath() + filepath},
	}
}
