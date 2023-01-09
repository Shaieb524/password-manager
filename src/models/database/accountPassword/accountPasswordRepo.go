package accountPassword

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func (repo *AccoutnPasswordRepo) CreateAccountPassword(accPassInput AccountPasswordInputDto) (*AccountPasswordInputDto, error) {
	var dbAccPass AccountPassword
	dbAccPass.Service = accPassInput.Service
	dbAccPass.Password = accPassInput.Password
	err := repo.db.Create(&dbAccPass).Error
	if err != nil {
		repo.logger.Error("Couldn't create service password into DB")
		return &AccountPasswordInputDto{}, err
	}
	return &accPassInput, nil
}

func (repo *AccoutnPasswordRepo) GetAllAccountsPasswords() (*[]AccountPassword, error) {
	var accountsPasswords *[]AccountPassword

	err := repo.db.Find(&accountsPasswords).Error
	if err != nil {
		repo.logger.Error("Couldn't get service passwords from DB")
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
		repo.logger.Error("Couldn't get service password from DB")
		return nil, err
	}
	return &ap, nil
}

func (repo *AccoutnPasswordRepo) EditAccountPassword(accPassInput AccountPasswordInputDto) (*AccountPasswordInputDto, error) {
	var ap AccountPassword
	if err := repo.db.First(&ap, "service=?", accPassInput.Service).Error; err != nil {
		repo.logger.Error("Couldn't find service password in DB")
		return nil, err
	}
	ap.Password = accPassInput.Password
	repo.db.Save(&ap)

	return &accPassInput, nil
}

func (repo *AccoutnPasswordRepo) DeleteByName(serviceName string) error {
	var ap AccountPassword
	if err := repo.db.First(&ap, "service=?", serviceName).Error; err != nil {
		repo.logger.Error("Couldn't find service password in DB")
		return err
	}
	repo.db.Delete(&ap)
	return nil
}

// DI
func NewAccPasswordRepoModule(logger *zap.Logger, db *gorm.DB) *AccoutnPasswordRepo {
	return &AccoutnPasswordRepo{
		logger: logger,
		db:     db,
	}
}
