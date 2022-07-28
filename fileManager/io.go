package fileManager

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"

	"github.com/kerrrusha/BTC-API/error"
	"github.com/kerrrusha/BTC-API/model"
)

func CreateEmptyEmailsJSON(filename string) int {
	file, err := os.Create(filename)
	error.CheckForError(err)

	emails := model.Emails{Emails: []string{}}
	emailsJSON, err := json.Marshal(emails)
	error.CheckForError(err)

	length, err := io.WriteString(file, string(emailsJSON))
	error.CheckForError(err)

	defer file.Close()

	return length
}

func WriteToFile(filename string, content []byte) int {
	file, err := os.Create(filename)
	error.CheckForError(err)

	length, err := io.WriteString(file, string(content))
	error.CheckForError(err)

	defer file.Close()

	return length
}

func ReadFile(filename string) []byte {
	databyte, err := ioutil.ReadFile(filename)
	if os.IsNotExist(err) {
		CreateEmptyEmailsJSON(filename)
		databyte, err = ioutil.ReadFile(filename)
		error.CheckForError(err)
	}
	return databyte
}
