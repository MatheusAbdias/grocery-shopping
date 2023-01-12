package domain

import (
	"github.com/google/uuid"
)

type User struct {
	Id       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Role     string    `json:"role"`
}

func NewUser(user UserInputDTO) (*User, []error) {
	if errs := Validate(user); len(errs) != 0 {
		return nil, errs
	}

	return &User{
		Id:       uuid.New(),
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Role:     user.Role,
	}, nil
}

func Validate(user UserInputDTO) []error {
	errs := []error{}

	validateRole(user.Role, &errs)
	ValidateEmail(user.Email, &errs)

	return errs
}
