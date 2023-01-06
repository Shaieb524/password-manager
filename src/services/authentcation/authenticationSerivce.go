package authentication

import (
	"password-manager/src/models/database/user"
)

type AuthenticationService struct {
	repo *user.AuthenticationRepo
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
func NewAuthenticationServiceModule(repo *user.AuthenticationRepo) *AuthenticationService {
	return &AuthenticationService{repo: repo}
}
