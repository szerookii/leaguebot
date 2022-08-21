package database

import (
	"github.com/szerookii/leaguebot/database/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Init() {
	db, err := gorm.Open(sqlite.Open("./database.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Summoner{})
}

func GetDB() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open("./database.db"), &gorm.Config{})
}
