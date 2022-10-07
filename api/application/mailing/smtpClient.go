package mailing

import (
	"github.com/kerrrusha/btc-api/logger"
	"net/smtp"

	"github.com/kerrrusha/btc-api/api/internal/config"
	"github.com/kerrrusha/btc-api/api/internal/customErrors"
)

type SmtpClient struct {
	identity string
	username string
	password string
	from     string
	host     string
	port     string

	auth smtp.Auth
}

func (s *SmtpClient) getAddress() string {
	return s.host + ":" + s.port
}

func (s *SmtpClient) getMessage(subject string, body string) []byte {
	return []byte("From: " + s.from + "\r\n" +
		"To: you\r\n" +
		"Subject: " + subject + "\r\n\r\n" +
		body + "\r\n")
}

func (s *SmtpClient) authorize() {
	s.auth = smtp.PlainAuth(s.identity, s.username, s.password, s.host)
	log := logger.CreateRabbitMQLogger()
	log.Info("SMTP was authorized successfully.")
}

func (s *SmtpClient) SendEmails(to []string, subject string, body string) *customErrors.SendMailError {
	if len(to) <= 0 {
		return nil
	}

	msg := s.getMessage(subject, body)
	err := smtp.SendMail(s.getAddress(), s.auth, s.from, to, msg)
	if err != nil {
		errMsg := "Failed to send emails"
		log := logger.CreateRabbitMQLogger()
		log.Error(errMsg + ": " + err.Error())
		return customErrors.CreateSendMailError(errMsg)
	}

	return nil
}

func createSmtpClient(identity string, username string, password string,
	from string, host string, port string) *SmtpClient {
	s := &SmtpClient{
		identity: identity,
		username: username,
		password: password,
		from:     from,
		host:     host,
		port:     port,
	}
	s.authorize()

	return s
}

func CreateSmtpEmailClient() EmailClient {
	var cl EmailClient
	cfg := config.GetConfig()
	cl = createSmtpClient(cfg.GetSmtpIdentity(), cfg.GetSmtpUsername(), cfg.GetSmtpPassword(),
		cfg.GetSmtpFrom(), cfg.GetSmtpHost(), cfg.GetSmtpPort())
	return cl
}
