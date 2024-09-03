package models

import (
	"gorm.io/gorm"
)

func Setup(db *gorm.DB) error {
	err := db.AutoMigrate(&PublicIpLog{})
	return err
}
