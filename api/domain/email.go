package domain

import "net/mail"

type Email struct {
	Email string `json:"email"`
}

func (email *Email) IsValid() bool {
	_, err := mail.ParseAddress(email.Email)
	return err == nil
}

type Emails struct {
	Emails []string `json:"emails"`
}

func (emails *Emails) ToList() []string {
	return emails.Emails
}
func (emails *Emails) Append(email Email) {
	emails.Emails = append(emails.Emails, email.Email)
}
func (emails *Emails) ToString() string {
	result := "["
	for i, emailStr := range emails.ToList() {
		sep := ""
		if i != len(emails.ToList())-1 {
			sep = ", "
		}
		result += emailStr + sep
	}
	result += "]"

	return result
}
