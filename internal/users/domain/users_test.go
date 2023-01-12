package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewUSer(t *testing.T) {
	user := UserInputDTO{
		Name:  "John Doe",
		Email: "test@test.com",
		Role:  "customer",
	}

	newUser, errs := NewUser(user)

	assert.Equal(t, errs, []error(nil))
	assert.Equal(t, newUser.Name, user.Name)
	assert.Equal(t, newUser.Email, user.Email)
	assert.Equal(t, newUser.Role, user.Role)

}

func TestFailCreateNewUserPassingInvalidRole(t *testing.T) {
	user := UserInputDTO{
		Name:  "John Doe",
		Email: "test@test.com",
		Role:  "invalid",
	}

	_, errs := NewUser(user)

	assert.Error(t, errs[0], "invalid role")
}

func TestFailCreateNewUserMissingPassingInvalidEmail(t *testing.T) {
	user := UserInputDTO{
		Name:  "John Doe",
		Email: "test.com",
		Role:  "customer",
	}

	_, errs := NewUser(user)

	assert.Error(t, errs[0], "invalid email")
}
