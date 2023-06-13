package main

import (
	"go-boiler-plate/config"
	"go-boiler-plate/server"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	activeProfile := os.Getenv("ACTIVE_PROFILE")
	conf, err := config.LoadConfig(activeProfile)
	if err != nil {
		log.Panic("Failed to load config ", err)
	}
	config.LoadLocale(&conf)
	// logger := logger.Init(&conf)

	dbStore := config.InitializeDB(&conf)

	server, err := server.Init(&conf, dbStore)
	if err != nil {
		log.Fatal("Failer get server", err)
	}
	err = server.Start(conf.HTTPServerAddress)
	if err != nil {
		log.Fatal("Unable to start server", err)
	}
}
