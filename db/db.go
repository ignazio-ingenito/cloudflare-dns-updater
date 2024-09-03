package db

import (
	"dnsupdater/models"
	"dnsupdater/types"
	"log"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const DB_NAME = "data.db"

type TPublicIpLogUpdate struct {
	Ip        string
	CreatedAt string
}

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(DB_NAME), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Panicln(err)
	}

	switch os.Getenv("LOG_LEVEL") {
	case "info":
		logger.Default.LogMode(logger.Info)
	case "warn":
		logger.Default.LogMode(logger.Warn)
	case "error":
		logger.Default.LogMode(logger.Error)
	case "silent":
		logger.Default.LogMode(logger.Silent)
	default:
		logger.Default.LogMode(logger.Warn)
	}

	return db, err
}

func PublicIpLogAll(sql *gorm.DB, limit int, reverse bool) []models.PublicIpLog {
	var ips []models.PublicIpLog

	order_text := "created_at ASC"
	if reverse {
		order_text = "created_at DESC"
	}
	sql.Order(order_text).Limit(limit).Find(&ips)

	return ips
}

func PublicIpLogUpdateAll(sql *gorm.DB, limit int, reverse bool) []models.PublicIpLog {
	var logs []TPublicIpLogUpdate

	order_text := "CreatedAt ASC"
	if reverse {
		order_text = "CreatedAt DESC"
	}
	sql.Model(&models.PublicIpLog{}).Group("ip").Select("ip Ip, max(created_at) CreatedAt").Order(order_text).Limit(limit).Find(&logs)

	var rows []models.PublicIpLog
	for _, l := range logs {
		ip := models.PublicIpLog{
			Ip: l.Ip,
		}
		parsedTime, _ := time.Parse("2006-01-02 15:04:05.999999-07:00", l.CreatedAt)
		ip.CreatedAt = parsedTime
		rows = append(rows, ip)
	}
	return rows
}

func PublicIpLogCreate(sql *gorm.DB) {
	ip := &types.PublicIp{}
	err := ip.Get()
	if err != nil {
		log.Printf("Ip not found %s", err)
		return
	}

	sql.Create(&models.PublicIpLog{
		Ip: ip.Ip,
	})
}
