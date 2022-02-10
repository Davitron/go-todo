package db

import (
	"fmt"
	"go-todo/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(dbConfig config.DatabaseConfigurations) *gorm.DB {
	conn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		dbConfig.DBUser,
		dbConfig.DBPassword,
		dbConfig.DBHost,
		dbConfig.DBPort,
		dbConfig.DBName)
	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		panic("Error connection to Database")
	}
	runMigration(db)
	return db
}
