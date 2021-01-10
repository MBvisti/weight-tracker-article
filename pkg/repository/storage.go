package repository

import (
	"database/sql"
	"errors"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"log"
	"path/filepath"
	"runtime"
	"weight-tracker/pkg/api"
)

type Storage interface {
	RunMigrations(connectionString string) error
	CreateUser(request api.NewUserRequest) error
	UpdateActivityLevel(res api.UpdateActivityLevelRequest) error
}

type storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) Storage {
	return &storage{db: db}
}

func (s *storage) RunMigrations(connectionString string) error {
	if connectionString == "" {
		return errors.New("repository: the connString was empty")
	}
	// get base path
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Join(filepath.Dir(b), "../..")

	migrationsPath := filepath.Join("file://", basePath, "/pkg/repository/migrations/")

	m, err := migrate.New(migrationsPath, connectionString)

	if err != nil {
		return err
	}

	err = m.Up()

	switch err {
	case errors.New("no change"):
		return nil
	}

	return nil
}

func (s *storage) CreateUser(request api.NewUserRequest) error {
	newUserStatement := `
		INSERT INTO "user" (name, age, height, sex, activity_level, email) 
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id;
		`

	var ID int
	err := s.db.QueryRow(newUserStatement, request.Name, request.Age, request.Height, request.Sex, request.ActivityLevel, request.Email).Scan(&ID)

	log.Printf("this is id: %v", ID)
	if err != nil {
		log.Printf("this was the error: %v", err.Error())
		return err
	}


	return nil
}

// TODO: to be implemented
func (s *storage) UpdateActivityLevel(res api.UpdateActivityLevelRequest) error {
	updateActivityLevelStatement := `INSERT INTO users`

	err := s.db.QueryRow(updateActivityLevelStatement, res.Email).Err()

	if err != nil {
		return err
	}

	return nil
}