package accountPassword

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AccountPassword struct {
	gorm.Model
	Service  string `gorm:"size:255;not null;unique" json:"service"`
	Password string `gorm:"size:255;not null;" json:"password"`
}

type AccountPasswordInputDto struct {
	Service  string `json:"service"`
	Password string `json:"password"`
}

type CachedAccountPassword struct {
	AccountPasswordInputDto
	ExpireAtTimestamp int64
}

type AccoutnPasswordRepo struct {
	logger *zap.Logger
	db     *gorm.DB
}
