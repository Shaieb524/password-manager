package accountPassword

import (
	"gorm.io/gorm"
)

func (repo *AccoutnPasswordRepo) CreateAccountPassword(accPassInput AccountPasswordInputDto) (*AccountPasswordInputDto, error) {
	var dbAccPass AccountPassword
	dbAccPass.Service = accPassInput.Service
	dbAccPass.Password = accPassInput.Password
	err := repo.db.Create(&dbAccPass).Error
	if err != nil {
		return &AccountPasswordInputDto{}, err
	}
	return &accPassInput, nil
}

func (repo *AccoutnPasswordRepo) GetAllAccountsPasswords() (*[]AccountPassword, error) {
	var accountsPasswords *[]AccountPassword

	err := repo.db.Find(&accountsPasswords).Error
	if err != nil {
		return &[]AccountPassword{}, err
	}
	return accountsPasswords, nil
}

func (repo *AccoutnPasswordRepo) GetByID(id string) (*AccountPassword, error) {
	var ap AccountPassword
	if err := repo.db.First(&ap, "id=?", id).Error; err != nil {
		return nil, err
	}
	return &ap, nil
}

func (repo *AccoutnPasswordRepo) GetByServiceName(serviceName string) (*AccountPassword, error) {
	var ap AccountPassword
	if err := repo.db.First(&ap, "service=?", serviceName).Error; err != nil {
		return nil, err
	}
	return &ap, nil
}

func (repo *AccoutnPasswordRepo) EditAccountPassword(accPassInput AccountPasswordInputDto) (*AccountPasswordInputDto, error) {
	var ap AccountPassword
	if err := repo.db.First(&ap, "service=?", accPassInput.Service).Error; err != nil {
		return nil, err
	}
	ap.Password = accPassInput.Password
	repo.db.Save(&ap)

	return &accPassInput, nil
}

func (repo *AccoutnPasswordRepo) DeleteByName(serviceName string) error {
	var ap AccountPassword
	if err := repo.db.First(&ap, "service=?", serviceName).Error; err != nil {
		return err
	}
	repo.db.Delete(&ap)
	return nil
}

// DI
func NewAccPasswordRepoModule(db *gorm.DB) *AccoutnPasswordRepo {
	return &AccoutnPasswordRepo{db: db}
}
