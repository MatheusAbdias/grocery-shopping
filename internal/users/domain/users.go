package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id       uuid.UUID `json:"id" db:"id"`
	Name     string    `json:"name" db:"name"`
	Email    string    `json:"email" db:"email"`
	Password string    `json:"password" db:"password"`
	Role     string    `json:"role" db:"role"`
	Created  time.Time `json:"created" db:"created"`
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
		Created:  time.Now(),
	}, nil
}

func Validate(user UserInputDTO) []error {
	errs := []error{}

	validateRole(user.Role, &errs)
	ValidateEmail(user.Email, &errs)

	return errs
}
