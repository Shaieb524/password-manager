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
	EditServicePassword(accPass accountPassword.AccountPassword)
	DeleteServicePassword(serviceName string)
}

func (apS *AccountPasswordService) CreateAccountPassword(accPass accountPassword.AccountPasswordInputDto) (*accountPassword.AccountPasswordInputDto, error) {
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

func (apS *AccountPasswordService) EditServicePassword(accPass accountPassword.AccountPasswordInputDto) (*accountPassword.AccountPasswordInputDto, error) {
	return apS.repo.EditAccountPassword(accPass)
}

func (apS *AccountPasswordService) DeleteServicePassword(serviceName string) error {
	return apS.repo.DeleteByName(serviceName)
}

// DI
func NewAccPasswordServiceModule(repo *accountPassword.AccoutnPasswordRepo) *AccountPasswordService {
	return &AccountPasswordService{repo: repo}
}
