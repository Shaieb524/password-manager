package accountPassword

import (
	"gorm.io/gorm"
)

type AccountPassword struct {
	gorm.Model
	Service  string `gorm:"size:255;not null;unique" json:"service"`
	Password string `gorm:"size:255;not null;" json:"password"`
}

type AccountPasswordInputDto struct {
	Service  string `gorm:"size:255;not null;unique" json:"service"`
	Password string `gorm:"size:255;not null;" json:"password"`
}

type AccoutnPasswordRepo struct {
	db *gorm.DB
}
