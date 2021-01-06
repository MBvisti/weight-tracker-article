package api

// UserService contains the methods of the user service
type UserService interface {
}

// User repository is what lets our service do db operations without knowing anything about the implementation
type UserRepository interface {
}

type userService struct {
	storage UserRepository
}

func NewUserService(userRepo UserRepository) UserService {
	return &userService{storage: userRepo}
}
