package models

import (
	"gorm.io/gorm"
)

type PublicIpLog struct {
	Ip string `gorm:"not null" json:"ip"`
	gorm.Model
}
