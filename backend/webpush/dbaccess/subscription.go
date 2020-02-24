package dbaccess

import (
	"github.com/jinzhu/gorm"
)

type Subscription struct {
	gorm.Model
	UserID   int    `gorm:"size:11"`
	Endpoint string `gorm:"size:2048"`
	P256dh   string `gorm:"size:255"`
	Auth     string `gorm:"size:255"`
}
