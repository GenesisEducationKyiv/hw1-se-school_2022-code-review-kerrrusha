package rest

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kerrrusha/BTC-API/api/internal/arrayUtils"
	"github.com/kerrrusha/BTC-API/api/internal/config"
	"github.com/kerrrusha/BTC-API/api/internal/errorUtils"
	"github.com/kerrrusha/BTC-API/api/internal/model"
	"github.com/kerrrusha/BTC-API/api/internal/model/dataStorage/fileStorage"
	"github.com/kerrrusha/BTC-API/api/internal/responseUtils"
)

func CreateEmptyEmailsJSON(filepath string) {
	emails := model.Emails{Emails: []string{}}
	emailsJSON, err := json.Marshal(emails)
	errorUtils.CheckForError(err)

	writer := fileStorage.CreateFileWriter(filepath)
	writer.Write(string(emailsJSON), false)
}

func ReadEmails(filepath string) model.Emails {
	var emails model.Emails

	if fileStorage.FileNotExist(filepath) || fileStorage.FileIsEmpty(filepath) {
		CreateEmptyEmailsJSON(filepath)
	}

	reader := fileStorage.CreateFileReader(filepath)
	fileBytes := reader.Read()
	err := json.Unmarshal(fileBytes, &emails)
	errorUtils.CheckForError(err)

	return emails
}

func WriteNewEmailToFile(filepath string, email string) {
	var emails model.Emails
	storage := fileStorage.CreateFileStorage(filepath)

	fileBytes := storage.Read()

	err := json.Unmarshal(fileBytes, &emails)
	errorUtils.CheckForError(err)

	emails.Emails = append(emails.Emails, email)

	emailsJSON, err := json.Marshal(emails)
	errorUtils.CheckForError(err)

	storage.Write(string(emailsJSON), true)
}

func Subscribe(w http.ResponseWriter, r *http.Request) {
	log.Println("subscribe endpoint")

	decoder := json.NewDecoder(r.Body)
	cfg := config.Get()

	var newEmail model.Email
	err := decoder.Decode(&newEmail)
	errorUtils.CheckForError(err)

	if arrayUtils.StringArraySearch(ReadEmails(cfg.Filepath).Emails, newEmail.Email) != -1 {
		responseUtils.SendResponse(
			w,
			model.ErrorResponse{Error: "Email was not subscribed: it already exists"},
			http.StatusConflict,
		)
		return
	}
	if !newEmail.IsValid() {
		responseUtils.SendResponse(
			w,
			model.ErrorResponse{Error: "Email is not correct. Please, enter valid email"},
			http.StatusConflict,
		)
		return
	}

	WriteNewEmailToFile(cfg.Filepath, newEmail.Email)

	responseUtils.SendResponse(
		w,
		model.SuccessResponse{Success: "Email was subscribed successfully"},
		http.StatusOK,
	)
}
