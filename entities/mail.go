package entities

import (
	"regexp"

	"github.com/Renan-Parise/codium-mail/errors"
)

type Email struct {
	Address string `json:"address"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

func (e *Email) Validate() error {
	if e.Address == "" {
		return errors.NewValidationError("address", "email address is required")
	}

	if !regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`).MatchString(e.Address) {
		return errors.NewValidationError("address", "invalid email address")
	}

	if e.Subject == "" {
		return errors.NewValidationError("subject", "email subject is required")
	}

	if e.Body == "" {
		return errors.NewValidationError("body", "email body is required")
	}

	return nil
}
