package main

import (
	"fmt"
	"go-todo/api"
	c "go-todo/config"
	"go-todo/core/db"
	"log"
)

func main() {
	config, _ := c.Init(".")
	port := config.Server.Port
	serverAddress := fmt.Sprintf("0.0.0.0:%s", port)
	DB := db.InitDB(config.Database)
	server := api.NewServer(DB, config)
	err := server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}
}
