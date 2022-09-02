package tests

import (
	"testing"

	"github.com/kerrrusha/BTC-API/fileManager"
	"github.com/stretchr/testify/assert"
)

const isEmptyErrorMsg = "File is not empty."

func TestFileIsEmpty(t *testing.T) {
	filename := "empty.json"

	assert.Equal(t, true,
		fileManager.FileIsEmpty(filename), isEmptyErrorMsg)
}

func TestFileIsNotEmpty(t *testing.T) {
	filename := "notEmpty.json"

	assert.Equal(t, true,
		fileManager.FileIsEmpty(filename), isEmptyErrorMsg)
}

func TestFileError(t *testing.T) {
	filename := "sfbdjbsdkf.json"
	assert.Panics(t, func() { fileManager.FileIsEmpty(filename) })
}
