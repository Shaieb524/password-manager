package accountPassword

import (
	"gorm.io/gorm"
)

type AccountPasswordModel struct {
	gorm.Model
	Service  string `gorm:"size:255;not null;unique" json:"service"`
	Password string `gorm:"size:255;not null;" json:"-"`
}

type AccoutnPasswordRepo struct {
	db *gorm.DB
}
