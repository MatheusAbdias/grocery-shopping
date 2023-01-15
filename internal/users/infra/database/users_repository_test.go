package database

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/MatheusAbdias/grocery-shopping/internal/users/domain"
	"github.com/MatheusAbdias/grocery-shopping/pkg/postgres"
	"github.com/MatheusAbdias/grocery-shopping/pkg/postgres/configs"
	"github.com/stretchr/testify/suite"
)

type UsersRepositoryTestSuite struct {
	suite.Suite
	Db *postgres.Database
}

func (suite *UsersRepositoryTestSuite) SetupSuite() {
	err := postgres.OpenConnection(&configs.DBConfig{
		Type:     "postgres",
		Host:     "localhost",
		DBPort:   "5432",
		User:     "grocery_user",
		Password: "password",
		Database: "grocery",
	})

	suite.NoError(err)

	suite.Db = postgres.GetDB()

	path, err := os.Executable()
	suite.NoError(err)

	basePath := filepath.Join(path, "../migrations")
	postgres.RunMigrations(suite.Db, basePath)
}

func (S *UsersRepositoryTestSuite) TearDownSuite() {
	S.Db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(UsersRepositoryTestSuite))
}

func (suite *UsersRepositoryTestSuite) TestGivenAnUser_WhenSave_ThenShouldSaveUser(t *testing.T) {

	user := domain.UserInputDTO{
		Name:  "John Doe",
		Email: "test@test.com",
		Role:  "customer",
	}
	newUser, errs := domain.NewUser(user)

	suite.Assert().Len(errs, 0)

	repository := NewUsersRepository(suite.Db)
	errs = append(errs, repository.Save(newUser))

	suite.Assert().Len(errs, 0)
}
