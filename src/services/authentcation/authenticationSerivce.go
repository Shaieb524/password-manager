package authentication

import (
	"password-manager/src/models/database/user"

	"go.uber.org/zap"
)

type AuthenticationService struct {
	logger *zap.Logger
	repo   *user.AuthenticationRepo
}

type authenticationService interface {
	ValidatePassword(existingUser *user.User, inputPassword string)
	FindUserByName(inputPassword string)
	RegisterUser(user *user.User)
}

func (authS *AuthenticationService) ValidatePassword(existingUser *user.User, inputPassword string) error {
	return authS.repo.ValidatePassword(existingUser, inputPassword)
}

func (authS *AuthenticationService) FindUserByName(inputUserName string) (*user.User, error) {
	return authS.repo.FindUserByUsername(inputUserName)
}

func (authS *AuthenticationService) RegisterUser(user *user.User) (*user.User, error) {
	return authS.repo.Save(user)
}

// DI
func NewAuthenticationServiceModule(logger *zap.Logger, repo *user.AuthenticationRepo) *AuthenticationService {
	return &AuthenticationService{
		logger: logger,
		repo:   repo,
	}
}
