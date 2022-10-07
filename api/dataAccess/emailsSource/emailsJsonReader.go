package emailsSource

import (
	"encoding/json"

	"github.com/kerrrusha/btc-api/api/dataAccess/storage"
	"github.com/kerrrusha/btc-api/api/dataAccess/storage/fileStorage"
	"github.com/kerrrusha/btc-api/api/domain"
	"github.com/kerrrusha/btc-api/api/internal/customErrors"
	"github.com/kerrrusha/btc-api/api/internal/utils"
)

type EmailsJsonReader struct {
	reader   storage.DataReadable
	jsonpath string
}

func (e *EmailsJsonReader) GetEmails() (*domain.Emails, *customErrors.CustomError) {
	var emails domain.Emails

	if utils.FileNotExist(e.jsonpath) || utils.FileIsEmpty(e.jsonpath) {
		return nil, customErrors.CreateFileIsBrokenError("Emails json file not exists or empty").CustomError
	}

	fileBytes := e.reader.Read()
	err := json.Unmarshal(fileBytes, &emails)
	if err != nil {
		return nil, customErrors.CreateJsonUnmarshalError("Json unmarshal failure: " + e.jsonpath).CustomError
	}

	return &emails, nil
}

func CreateJsonEmailsReader(jsonpath string) *EmailsJsonReader {
	fileReader := fileStorage.CreateFileReader(jsonpath)
	return &EmailsJsonReader{reader: fileReader, jsonpath: jsonpath}
}
