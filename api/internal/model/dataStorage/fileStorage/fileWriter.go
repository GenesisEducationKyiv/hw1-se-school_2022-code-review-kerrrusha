package fileStorage

import (
	"io"

	"github.com/kerrrusha/BTC-API/api/internal/errorUtils"
	"github.com/kerrrusha/BTC-API/api/internal/model/dataStorage"
)

type fileWriter struct {
	*dataStorage.FileAccessable
}

func (writer *fileWriter) Write(content string, append bool) int {
	file := writer.AccessFileWrite()

	length, err := io.WriteString(file, content)
	errorUtils.CheckForError(err)

	defer file.Close()

	return length
}

func CreateFileWriter(filepath string) *fileWriter {
	return &fileWriter{
		FileAccessable: &dataStorage.FileAccessable{Path: GetProjPath() + filepath},
	}
}
