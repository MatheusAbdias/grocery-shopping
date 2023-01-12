package domain

import (
	"errors"
	"net/mail"
)

func validateRole(role string, errs *[]error) {
	if !(role == "admin" || role == "customer") {
		*errs = append(*errs, errors.New("invalid role"))

	}
}

func ValidateEmail(email string, errs *[]error) {
	_, err := mail.ParseAddress(email)
	if err != nil {
		*errs = append(*errs, errors.New("invalid email"))
	}
}
