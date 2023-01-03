package services

import (
	"password-manager/src/models/database/accountPassword"
)

type AccountPasswordService struct {
	repo accountPassword.AccoutnPasswordRepo
}

type accountPasswordService interface {
	CreatAccountPassword(accPass accountPassword.AccountPasswordModel)
	GetAppPasswordById()
}

func (apS *AccountPasswordService) CreateAccountPassword(accPass accountPassword.AccountPasswordModel) (*accountPassword.AccountPasswordModel, error) {
	return apS.repo.CreateAccountPassword(accPass)
}

func (apS *AccountPasswordService) GetAppPasswordById(id string) (*accountPassword.AccountPasswordModel, error) {
	return apS.repo.GetByID(id)
}