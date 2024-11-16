package agent

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/Renan-Parise/mail/entities"
	"github.com/Renan-Parise/mail/errors"
)

func SendEmail(email entities.Email) error {
	gmailUser := os.Getenv("GMAIL_USERNAME")
	gmailPass := os.Getenv("GMAIL_PASSWORD")

	if gmailUser == "" || gmailPass == "" {
		return errors.NewValidationError("email", "Gmail credentials not set")
	}

	auth := smtp.PlainAuth("", gmailUser, gmailPass, "smtp.gmail.com")

	from := gmailUser
	to := []string{email.Address}
	subject := email.Subject
	body := email.Body

	message := []byte(fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s",
		from, email.Address, subject, body))

	err := smtp.SendMail("smtp.gmail.com:587", auth, from, to, message)
	if err != nil {
		errors.NewServiceError("Failed to send email: " + err.Error())
	}

	return nil
}
