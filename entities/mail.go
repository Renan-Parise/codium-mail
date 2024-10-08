package entities

import (
	"errors"
	"regexp"
)

type Email struct {
	Address string `json:"address"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

func (e *Email) Validate() error {
	if e.Address == "" {
		return errors.New("email address is required")
	}

	if !regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`).MatchString(e.Address) {
		return errors.New("invalid email address")
	}

	if e.Subject == "" {
		return errors.New("email subject is required")
	}

	if e.Body == "" {
		return errors.New("email body is required")
	}

	return nil
}
