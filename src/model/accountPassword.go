package model

import (
	"password-manager/src/config/database"

	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AccountPassword struct {
	gorm.Model
	Service  string `gorm:"size:255;not null;unique" json:"service"`
	Password string `gorm:"size:255;not null;" json:"-"`
}

func (ap *AccountPassword) Save() (*AccountPassword, error) {
	err := database.Database.Create(&ap).Error
	if err != nil {
		return &AccountPassword{}, err
	}
	return ap, nil
}

func (ap *AccountPassword) BeforeSave() error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(ap.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	ap.Password = string(passwordHash)
	ap.Service = html.EscapeString(strings.TrimSpace(ap.Service))
	return nil
}
