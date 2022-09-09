package fileStorage

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/kerrrusha/BTC-API/api/internal/config"
	"github.com/kerrrusha/BTC-API/api/internal/errorUtils"
)

func FileNotExist(filepath string) bool {
	_, err := os.Stat(filepath)
	return errors.Is(err, os.ErrNotExist)
}

func FileIsEmpty(filepath string) bool {
	fileBytes := CreateFileReader(filepath).Read()
	return len(fileBytes) <= 0
}

func createFile(filepath string) {
	_, err := os.Create(filepath)
	errorUtils.CheckForError(err)
}

func GetGoSrcPath() string {
	ex, err := os.Getwd()
	errorUtils.CheckForError(err)
	return filepath.Dir(ex)
}

func GetProjPath() string {
	sep := "\\"
	return GetGoSrcPath() + sep + config.Get().ProjectName + sep
}
