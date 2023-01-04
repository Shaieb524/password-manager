package services

import (
	"password-manager/src/models/database/accountPassword"
)

type AccountPasswordService struct {
	repo *accountPassword.AccoutnPasswordRepo
}

type accountPasswordService interface {
	CreatAccountPassword(accPass accountPassword.AccountPassword)
	GetAllAccountsPasswords()
	GetAppPasswordById(id string)
	GetAppPasswordByServiceName(serviceName string)
}

func (apS *AccountPasswordService) CreateAccountPassword(accPass accountPassword.AccountPassword) (*accountPassword.AccountPassword, error) {
	return apS.repo.CreateAccountPassword(accPass)
}

func (apS *AccountPasswordService) GetAllAccountsPasswords() (*[]accountPassword.AccountPassword, error) {
	return apS.repo.GetAllAccountsPasswords()
}

func (apS *AccountPasswordService) GetAppPasswordById(id string) (*accountPassword.AccountPassword, error) {
	return apS.repo.GetByID(id)
}

func (apS *AccountPasswordService) GetAppPasswordByServiceName(serviceName string) (*accountPassword.AccountPassword, error) {
	return apS.repo.GetByServiceName(serviceName)
}

func ProvideModuleforDI(repo *accountPassword.AccoutnPasswordRepo) *AccountPasswordService {
	return &AccountPasswordService{repo: repo}
}
