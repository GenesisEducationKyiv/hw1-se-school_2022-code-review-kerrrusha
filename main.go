package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"time"
)

const (
	FILENAME    = "./emails.json"
	URL_GET_BTC = "https://rest.coinapi.io/v1/exchangerate/BTC/UAH?apikey=735B916A-29E3-49D7-BB21-5142DF49DAAC"
)

type RateAnswer struct {
	Time           string
	Asset_id_base  string
	Asset_id_quote string
	Rate           float64
}
type RateValue struct {
	Rate uint32 `json:"rate"`
}

type Email struct {
	Email string `json:"email"`
}
type Emails struct {
	Emails []string `json:"emails"`
}

type ErrorAnswer struct {
	Error string
}
type SuccessAnswer struct {
	Success string
}

func checkForError(err error) {
	if err != nil {
		panic(err)
	}
}

func sendErrorResponse(w http.ResponseWriter, msg string, code int) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	errorAnswer := ErrorAnswer{msg}
	json.NewEncoder(w).Encode(errorAnswer)
}

func sendSuccessResponse(w http.ResponseWriter, msg string, code int) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	successAnswer := SuccessAnswer{msg}
	json.NewEncoder(w).Encode(successAnswer)
}

func getJson(url string) []byte {
	client := http.Client{Timeout: time.Second * 2}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	checkForError(err)

	res, err := client.Do(req)
	checkForError(err)

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	checkForError(readErr)

	return body
}

func getBitcoinPriceUAH() (int, string) {
	jsonAnswer := getJson(URL_GET_BTC)

	var rateAnswer RateAnswer
	err := json.Unmarshal(jsonAnswer, &rateAnswer)
	checkForError(err)

	var errorAnswer ErrorAnswer
	err = json.Unmarshal(jsonAnswer, &errorAnswer)
	checkForError(err)

	if len(errorAnswer.Error) > 0 {
		return -1, errorAnswer.Error
	}

	rateUAH := int(rateAnswer.Rate)

	return rateUAH, ""
}

func rate(w http.ResponseWriter, r *http.Request) {
	log.Println("rate endpoint")

	result, errorMsg := getBitcoinPriceUAH()

	if len(errorMsg) > 0 {
		sendErrorResponse(w, errorMsg, http.StatusBadRequest)
		return
	}

	rateUAH := RateValue{uint32(result)}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(rateUAH)
}

func createEmptyEmailsJSON(filename string) int {
	file, err := os.Create(filename)
	checkForError(err)

	emails := Emails{[]string{}}
	emailsJSON, err := json.Marshal(emails)
	checkForError(err)

	length, err := io.WriteString(file, string(emailsJSON))
	checkForError(err)

	defer file.Close()

	return length
}

func writeToFile(filename string, content []byte) int {
	file, err := os.Create(filename)
	checkForError(err)

	length, err := io.WriteString(file, string(content))
	checkForError(err)

	defer file.Close()

	return length
}

func readFile(filename string) []byte {
	databyte, err := ioutil.ReadFile(filename)
	checkForError(err)
	return databyte
}

func readEmails(filename string) Emails {
	var emails Emails
	fileBytes := readFile(filename)

	if len(fileBytes) <= 0 {
		createEmptyEmailsJSON(filename)
		fileBytes = readFile(filename)
	}

	err := json.Unmarshal(fileBytes, &emails)
	checkForError(err)

	return emails
}

func indexOfEmail(filename string, email string) int {
	emails := readEmails(filename)

	for index, element := range emails.Emails {
		if element == email {
			return index
		}
	}

	return -1
}

func writeNewEmailToFile(filename string, email string) {
	var emails Emails
	fileBytes := readFile(filename)

	err := json.Unmarshal(fileBytes, &emails)
	checkForError(err)

	emails.Emails = append(emails.Emails, email)

	emailsJSON, err := json.Marshal(emails)
	checkForError(err)

	writeToFile(filename, emailsJSON)
}

func subscribeNewEmail(w http.ResponseWriter, r *http.Request) {
	log.Println("subscribe endpoint")
	decoder := json.NewDecoder(r.Body)

	var newEmail Email
	err := decoder.Decode(&newEmail)
	checkForError(err)

	log.Println(newEmail.Email)

	if indexOfEmail(FILENAME, newEmail.Email) != -1 {
		sendErrorResponse(w, "Email was not subscribed: it already exists", http.StatusConflict)
		return
	}

	writeNewEmailToFile(FILENAME, newEmail.Email)

	sendSuccessResponse(w, "Email was subscribed successfully", http.StatusOK)
}

func sendEmails(to []string, subject string, body string) {
	const (
		FROM     = "smtp8317@gmail.com"
		USERNAME = "smtp8317@gmail.com"
		PASSWORD = "khtihhqqywqrryan"
		PORT     = "587"
		HOST     = "smtp.gmail.com"
		ADDRESS  = HOST + ":" + PORT
	)

	if len(to) <= 0 {
		return
	}

	auth := smtp.PlainAuth("", USERNAME, PASSWORD, HOST)
	log.Println("SMTP was authorized successfully.")

	msg := []byte("From: " + FROM + "\r\n" +
		"To: you\r\n" +
		"Subject: " + subject + "\r\n\r\n" +
		body + "\r\n")
	err := smtp.SendMail(ADDRESS, auth, FROM, to, msg)
	checkForError(err)

	log.Println("Emails was sent successfully via SMTP '" + HOST + "' host.")
}

func sendBTCRateMails(w http.ResponseWriter, r *http.Request) {
	emails := readEmails(FILENAME)

	result, errorMsg := getBitcoinPriceUAH()

	if len(errorMsg) > 0 {
		sendErrorResponse(w, errorMsg, http.StatusBadRequest)
		return
	}

	subject := "BTC/UAH"
	body := fmt.Sprintf("%d", result)

	sendEmails(emails.Emails, subject, body)

	sendSuccessResponse(w, "Emails was sent successfully!", http.StatusOK)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	log.Println(w, "homePage endpoint")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/rate/", rate)
	http.HandleFunc("/subscribe/", subscribeNewEmail)
	http.HandleFunc("/sendEmails/", sendBTCRateMails)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}

func main() {
	handleRequests()
}
