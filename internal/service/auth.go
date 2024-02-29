package service

import (
	"GoCodeEval/internal/model"
	"GoCodeEval/pkg/utils"
)

// AuthService handles the business logic for authentication.
type AuthService struct {
	// Here you would have dependencies like your user repository
}

// NewAuthService creates a new instance of AuthService.
func NewAuthService() *AuthService {
	return &AuthService{
		// Here you would initialize dependencies
	}
}

// Register registers a new user.
func (svc *AuthService) Register(req *model.RegisterReq) (*model.UserResponse, error) {
	// Here you would check if the email is already in use, hash the password, etc.
	// This is where your business logic would be implemented. For now, we'll just simulate the user registration.

	// Simulate user creation with unique identifier
	uuid, err := utils.NewUUID()
	if err != nil {
		return nil, err
	}

	user := &model.User{
		ID:       uuid,
		Email:    req.Email,
		Username: req.Username,
		// The password would normally be hashed
		Password: req.Password,
	}

	// Simulate saving the user to the database
	// ...

	// Prepare the response
	userResp := &model.UserResponse{
		ID:       user.ID,
		Username: user.Username,
	}

	return userResp, nil
}
