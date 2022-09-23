package utils

import (
	"errors"
	"os"
	"path/filepath"
)

func FileNotExist(filepath string) bool {
	_, err := os.Stat(filepath)
	return errors.Is(err, os.ErrNotExist)
}

func FileIsEmpty(filepath string) bool {
	fileStat, err := os.Stat(filepath)
	if err != nil {
		panic(err)
	}
	return fileStat.Size() == 0
}

func GetGoSrcPath() string {
	ex, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return filepath.Dir(ex)
}

func GetProjPath(projectName string) string {
	sep := "\\"
	return GetGoSrcPath() + sep + projectName + sep
}
