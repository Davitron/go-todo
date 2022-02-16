package models

import (
	"database/sql"
	"fmt"
	"go-todo/pkg/settings"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB
var sqlDB *sql.DB

type Model struct {
	ID int `gorm:"primary_key" json:"id"`
}

func Setup() {
	var err error
	fmt.Println(settings.AppSettings)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		settings.AppSettings.Database.DBHost,
		settings.AppSettings.Database.DBUser,
		settings.AppSettings.Database.DBPassword,
		settings.AppSettings.Database.DBName,
		settings.AppSettings.Database.DBPort,
	)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connection to databse, %s", err)
	}

	sqlDB, err = db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
}

func CloseDB() {
	defer sqlDB.Close()
}
