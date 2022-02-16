package settings

import (
	"github.com/spf13/viper"
	"log"
)

type Server struct {
	Port    string
	RunMode string
}

type Database struct {
	DBName     string
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
}

type JWT struct {
	SecretKey string
}

type Settings struct {
	Server   Server
	Database Database
	JWT      JWT
}

var AppSettings = &Settings{}

var envMaps = map[string]string{
	"server.port":         "PORT",
	"database.dbname":     "DB_NAME",
	"database.dbhost":     "DB_HOST",
	"database.dbpassword": "DB_PASSWORD",
	"database.dbuser":     "DB_USER",
}

func InitSettings() {
	var err error
	viper.SetConfigName("config")
	viper.AddConfigPath("../conf/")
	viper.SetConfigType("yml")
	viper.AutomaticEnv()

	for key, value := range envMaps {
		log.Printf("Binding %s to config", value)
		err = viper.BindEnv(key, value)
		if err != nil {
			log.Fatalf("Error binding %s to config", value)
			return
		}
	}

	if err = viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
		return
	}

	viper.SetDefault("server.runmode", "debug")
	viper.SetDefault("server.port", 8080)
	viper.SetDefault("database.dbhost", "localhost")
	viper.SetDefault("database.dbport", "5432")

	err = viper.Unmarshal(&AppSettings)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
		return
	}
}
