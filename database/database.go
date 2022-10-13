package database

import (
	"github.com/szerookii/leaguebot/database/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Init() {
	db, err := gorm.Open(sqlite.Open("./database.db"), &gorm.Config{ Logger: logger.Default.LogMode(logger.Silent) })

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Summoner{})

	dbConn, err := db.DB()
	if err != nil {
		panic("failed to get database")
	}

	defer dbConn.Close()
}

func GetDB() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open("./database.db"), &gorm.Config{ Logger: logger.Default.LogMode(logger.Silent) })
}
