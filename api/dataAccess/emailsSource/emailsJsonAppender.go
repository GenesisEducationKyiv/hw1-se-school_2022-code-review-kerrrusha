package emailsSource

import (
	"encoding/json"

	"github.com/kerrrusha/btc-api/api/dataAccess/storage/fileStorage"
	"github.com/kerrrusha/btc-api/api/domain"
	"github.com/kerrrusha/btc-api/api/internal/customErrors"
)

type EmailsJsonAppender struct {
	storage  *fileStorage.FileStorage
	jsonpath string
}

func (e *EmailsJsonAppender) Append(email domain.Email) *customErrors.CustomError {
	var emails domain.Emails
	fileBytes := e.storage.Read()

	err := json.Unmarshal(fileBytes, &emails)
	if err != nil {
		return customErrors.CreateJsonUnmarshalError("Json unmarshal failure: " + e.jsonpath).CustomError
	}

	emails.Append(email)

	emailsJSON, err := json.Marshal(emails)
	if err != nil {
		return customErrors.CreateJsonUnmarshalError("Json marshal failure: " + emails.ToString()).CustomError
	}

	e.storage.Write(string(emailsJSON), true)

	return nil
}

func CreateJsonEmailsAppender(jsonpath string) *EmailsJsonAppender {
	st := fileStorage.CreateFileStorage(jsonpath)
	return &EmailsJsonAppender{storage: st, jsonpath: jsonpath}
}
