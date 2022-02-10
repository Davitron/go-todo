package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func Init(configPath string) (config *Configurations, err error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(configPath)
	viper.AutomaticEnv()
	viper.SetConfigType("yml")
	viper.AutomaticEnv()

	f := func(key string, value string) bool {
		log.Printf("Binding %s to config", value)
		err = viper.BindEnv(key, value)
		if err != nil {
			log.Fatalf("Error binding %s to config", value)
		}
		return err == nil
	}

	failed := f("server.port", "PORT") &&
		f("database.dbname", "DB_NAME") &&
		f("database.dbhost", "DB_HOST") &&
		f("database.dbuser", "DB_USER") &&
		f("database.dbpassword", "DB_PASSWORD")

	if !failed {
		log.Fatalf("Error binding env vars, %s", err)
		return
	}

	var configuration Configurations

	if err = viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
		return
	}

	viper.SetDefault("server.runmode", "debug")
	viper.SetDefault("server.port", 8080)
	viper.SetDefault("database.dbhost", "localhost")
	viper.SetDefault("database.dbport", "5432")

	err = viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
		return
	}

	fmt.Println(&configuration)
	return &configuration, nil
}
