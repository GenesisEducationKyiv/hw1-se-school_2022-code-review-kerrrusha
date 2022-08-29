package service

import (
	"encoding/json"
	"log"
	"net/http"
	"net/mail"
	"strings"

	"github.com/kerrrusha/BTC-API/config"
	"github.com/kerrrusha/BTC-API/error"
	"github.com/kerrrusha/BTC-API/fileManager"
	"github.com/kerrrusha/BTC-API/model"
	"github.com/kerrrusha/BTC-API/response"
)

func readEmails(filename string) model.Emails {
	var emails model.Emails
	fileBytes := fileManager.ReadFile(filename)

	if len(fileBytes) <= 0 {
		fileManager.CreateEmptyEmailsJSON(filename)
		fileBytes = fileManager.ReadFile(filename)
	}

	err := json.Unmarshal(fileBytes, &emails)
	error.CheckForError(err)

	return emails
}

func indexOfEmail(filename string, email string) int {
	emails := readEmails(filename)

	for index, element := range emails.Emails {
		if strings.EqualFold(element, email) {
			return index
		}
	}

	return -1
}

func emailIsValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func writeNewEmailToFile(filename string, email string) {
	var emails model.Emails
	fileBytes := fileManager.ReadFile(filename)

	err := json.Unmarshal(fileBytes, &emails)
	error.CheckForError(err)

	emails.Emails = append(emails.Emails, email)

	emailsJSON, err := json.Marshal(emails)
	error.CheckForError(err)

	fileManager.WriteToFile(filename, emailsJSON)
}

func SubscribeNewEmail(w http.ResponseWriter, r *http.Request) {
	log.Println("subscribe endpoint")
	decoder := json.NewDecoder(r.Body)

	var newEmail model.Email
	err := decoder.Decode(&newEmail)
	error.CheckForError(err)

	log.Println(newEmail.Email)

	if indexOfEmail(config.FILENAME, newEmail.Email) != -1 {
		response.SendErrorResponse(w, "Email was not subscribed: it already exists", http.StatusConflict)
		return
	}
	if !emailIsValid(newEmail.Email) {
		response.SendErrorResponse(w, "Email is not correct. Please, enter valid email", http.StatusConflict)
		return
	}

	writeNewEmailToFile(config.FILENAME, newEmail.Email)

	response.SendSuccessResponse(w, "Email was subscribed successfully", http.StatusOK)
}
