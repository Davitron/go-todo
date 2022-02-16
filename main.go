package main

import (
	"fmt"
	"go-todo/models"
	"go-todo/pkg/settings"
	"go-todo/routers"
	"log"
	"net/http"
)

func init() {
	fmt.Println("here")
	settings.InitSettings()
	models.Setup()
}

func main() {
	fmt.Println(settings.AppSettings)
	endPoint := fmt.Sprintf(":%s", settings.AppSettings.Server.Port)
	fmt.Println(endPoint)
	routersInit := routers.InitRourter()
	s := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    0,
		WriteTimeout:   0,
		MaxHeaderBytes: 0,
	}
	log.Printf("[info] start http server listening on %s", endPoint)
	log.Fatalf("Error starting server:, %s ", s.ListenAndServe())
}
