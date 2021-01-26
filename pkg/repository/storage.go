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
	CreateWeightEntry(request api.Weight) error
	GetUser(userID int) (api.User, error)
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
		INSERT INTO "user" (name, age, height, sex, activity_level, email, weight_goal) 
		VALUES ($1, $2, $3, $4, $5, $6, $7);
		`

	err := s.db.QueryRow(newUserStatement, request.Name, request.Age, request.Height, request.Sex, request.ActivityLevel, request.Email, request.WeightGoal).Err()

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

func (s *storage) CreateWeightEntry(request api.Weight) error {
	newWeightStatement := `
		INSERT INTO weight (weight, user_id, bmr, daily_caloric_intake) 
		VALUES ($1, $2, $3, $4)
		RETURNING id;
		`

	var ID int
	err := s.db.QueryRow(newWeightStatement, request.Weight, request.UserID, request.BMR, request.DailyCaloricIntake).Scan(&ID)

	if err != nil {
		log.Printf("this was the error: %v", err.Error())
		return err
	}

	return nil
}

func (s *storage) GetUser(userID int) (api.User, error) {
	getUserStatement := `
		SELECT id, name, age, height, sex, activity_level, email, weight_goal FROM "user" 
		where id=$1;		
		`

	var user api.User
	err := s.db.QueryRow(getUserStatement, userID).Scan(&user.ID, &user.Name, &user.Age, &user.Height, &user.Sex, &user.ActivityLevel, &user.Email, &user.WeightGoal)

	if err != nil {
		log.Printf("this was the error: %v", err.Error())
		return api.User{}, err
	}

	return user, nil
}
