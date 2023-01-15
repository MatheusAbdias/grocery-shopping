package database

import (
	"github.com/MatheusAbdias/grocery-shopping/internal/users/domain"
	"github.com/MatheusAbdias/grocery-shopping/pkg/postgres"
)

type UsersRepository struct {
	Db *postgres.Database
}

func NewUsersRepository(db *postgres.Database) *UsersRepository {

	return &UsersRepository{Db: db}
}

func (r *UsersRepository) Save(user *domain.User) error {
	query := `INSERT INTO users (user_id, name, email, password, created) VALUES (?, ?, ?, ?, ?)`

	return postgres.Insert(r.Db.Db, query, *user)
}
