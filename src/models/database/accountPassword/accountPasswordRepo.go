package accountPassword

import (
	"gorm.io/gorm"
)

// "password-manager/src/providers/database"

// "html"
// "strings"

// "golang.org/x/crypto/bcrypt"

func ProvideAccountPasswordRepo(DB *gorm.DB) AccoutnPasswordRepo {
	return AccoutnPasswordRepo{db: DB}
}

func (repo *AccoutnPasswordRepo) CreateAccountPassword(accPass AccountPassword) (*AccountPassword, error) {
	err := repo.db.Create(&accPass).Error
	if err != nil {
		return &AccountPassword{}, err
	}
	return &accPass, nil
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

func (repo *AccoutnPasswordRepo) EditAccountPassword(accPass AccountPassword) (*AccountPassword, error) {
	var ap AccountPassword
	if err := repo.db.First(&ap, "service=?", accPass.Service).Error; err != nil {
		return nil, err
	}
	ap.Password = accPass.Password
	repo.db.Save(&ap)

	return &ap, nil
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
func ProvideModuleforDI(db *gorm.DB) *AccoutnPasswordRepo {
	return &AccoutnPasswordRepo{db: db}
}
