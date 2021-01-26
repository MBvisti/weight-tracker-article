package api

import (
	"errors"
	"strings"
)

// UserService contains the methods of the user service
type UserService interface {
	New(user NewUserRequest) error
	UpdateActivityLevel(res UpdateActivityLevelRequest) error
}

// User repository is what lets our service do db operations without knowing anything about the implementation
type UserRepository interface {
	CreateUser(NewUserRequest) error
	UpdateActivityLevel(UpdateActivityLevelRequest) error
}

type userService struct {
	storage UserRepository
}

func NewUserService(userRepo UserRepository) UserService {
	return &userService{storage: userRepo}
}

// TODO: note that we are not checking if the user already exists - we probably should not the article will be too long
func (u *userService) New(user NewUserRequest) error {
	// do some basic validations
	if user.Email == "" {
		return errors.New("user service - email required")
	}

	if user.Name == "" {
		return errors.New("user service - name required")
	}

	if user.WeightGoal == "" {
		return errors.New("user service - weight goal required")
	}

	// do some basic normalisation
	user.Name = strings.ToLower(user.Name)
	user.Email = strings.TrimSpace(user.Email)

	err := u.storage.CreateUser(user)

	if err != nil {
		return err
	}

	return nil
}

// TODO: activity level is loosely defined here - should be a database set value
func (u *userService) UpdateActivityLevel(res UpdateActivityLevelRequest) error {
	if res.Email == "" {
		return errors.New("user service - no email provided")
	}

	err := u.storage.UpdateActivityLevel(res)

	if err != nil {
		return err
	}

	return nil
}
