package postgres

import (
	"database/sql"
	"fmt"

	"github.com/MatheusAbdias/grocery-shopping/pkg/postgres/configs"
	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

type Database struct {
	Db         *sql.DB
	ConnString string
}

var db *Database

func (d *Database) Close() {
	d.Db.Close()
}

func (d *Database) GetConnString() string {
	return d.ConnString
}

func OpenConnection(conf *configs.DBConfig) error {
	conn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		conf.Host, conf.User, conf.Password, conf.Database, conf.DBPort)

	database, err := sql.Open("postgres", conn)
	if err != nil {
		return err
	}

	db = &Database{
		Db:         database,
		ConnString: conn,
	}
	return nil
}

func GetDB() *Database {
	return db
}

func Insert[T interface{}](db *sql.DB, query string, args T) error {
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(args)
	if err != nil {
		return err
	}

	return nil
}

func Select[T interface{}](db *sql.DB, query string, instance T) (*T, error) {
	row := db.QueryRow(query)
	err := row.Scan(instance)
	if err != nil {
		return nil, fmt.Errorf("could not find #{err}")
	}

	return &instance, nil
}

func RunMigrations(db *Database, path string) error {
	m, err := migrate.New(path, db.GetConnString())
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}
