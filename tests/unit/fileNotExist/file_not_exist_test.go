package tests

import (
	"testing"

	"github.com/kerrrusha/BTC-API/fileManager"
	"github.com/stretchr/testify/assert"
)

const notExistMsg = "FileNotExist returned incorrect value"

func TestFileNotExist(t *testing.T) {
	filename := "nkgdfjbdf.json"

	assert.Equal(t, true,
		fileManager.FileNotExist(filename), notExistMsg)
}

func TestFileExist(t *testing.T) {
	filename := "exist.json"

	assert.Equal(t, false,
		fileManager.FileNotExist(filename), notExistMsg)
}
