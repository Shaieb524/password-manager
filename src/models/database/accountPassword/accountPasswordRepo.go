package accountPassword

import (
	"fmt"

	"gorm.io/gorm"
)

// "password-manager/src/providers/database"

// "html"
// "strings"

// "golang.org/x/crypto/bcrypt"

func ProvideAccountPasswordRepo(DB *gorm.DB) AccoutnPasswordRepo {
	return AccoutnPasswordRepo{db: DB}
}

func (repo *AccoutnPasswordRepo) CreateAccountPassword(accPass AccountPasswordModel) (*AccountPasswordModel, error) {
	fmt.Println("repo : ", repo)
	fmt.Println("accPass : ", accPass)
	err := repo.db.Create(&accPass).Error
	if err != nil {
		return &AccountPasswordModel{}, err
	}
	return &accPass, nil
}

func (repo *AccoutnPasswordRepo) GetByID(id string) (*AccountPasswordModel, error) {
	var ap AccountPasswordModel
	if err := repo.db.First(&ap, "id=?", id).Error; err != nil {
		return nil, err
	}
	return &ap, nil
}

// func (ap AccountPassword) BeforeSave() error {
// 	passwordHash, err := bcrypt.GenerateFromPassword([]byte(ap.Password), bcrypt.DefaultCost)
// 	if err != nil {
// 		return err
// 	}
// 	ap.Password = string(passwordHash)
// 	ap.Service = html.EscapeString(strings.TrimSpace(ap.Service))
// 	return nil
// }

func ProvideModuleforDI(db *gorm.DB) *AccoutnPasswordRepo {
	return &AccoutnPasswordRepo{db: db}
}
