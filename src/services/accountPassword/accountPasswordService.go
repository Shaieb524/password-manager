package services

import (
	"fmt"
	"password-manager/src/models/database/accountPassword"
	appCache "password-manager/src/providers/appCache"
)

type AccountPasswordService struct {
	repo   *accountPassword.AccoutnPasswordRepo
	LCache *appCache.LocalCache
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
	// apS.AppCache.Set(accPass.Service,
	// 	accountPassword.AccountPasswordInputDto{Service: accPass.Service, Password: accPass.Password}, 5*time.Minute)
	apS.LCache.Update(accPass, 10)
	return apS.repo.CreateAccountPassword(accPass)
}

func (apS *AccountPasswordService) GetAllAccountsPasswords() (*[]accountPassword.AccountPassword, error) {
	return apS.repo.GetAllAccountsPasswords()
}

func (apS *AccountPasswordService) GetAppPasswordById(id string) (*accountPassword.AccountPassword, error) {
	return apS.repo.GetByID(id)
}

func (apS *AccountPasswordService) GetAppPasswordByServiceName(serviceName string) (interface{}, error) {
	cachedAccPassword, err := apS.LCache.Read(serviceName)
	if err != nil {
		fmt.Println("error getting from cache")
	}

	if &cachedAccPassword != nil {
		return &cachedAccPassword, nil
	} else {
		return apS.repo.GetByServiceName(serviceName)
	}
}

func (apS *AccountPasswordService) EditServicePassword(accPass accountPassword.AccountPasswordInputDto) (*accountPassword.AccountPasswordInputDto, error) {
	return apS.repo.EditAccountPassword(accPass)
}

func (apS *AccountPasswordService) DeleteServicePassword(serviceName string) error {
	return apS.repo.DeleteByName(serviceName)
}

// DI
func NewAccPasswordServiceModule(repo *accountPassword.AccoutnPasswordRepo, lCache *appCache.LocalCache) *AccountPasswordService {
	return &AccountPasswordService{
		repo:   repo,
		LCache: lCache,
	}
}
