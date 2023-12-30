package main

import (
	"log"

	"app/internal/config"
)

func main() {
	log.Print("Configuring app intializing")
	config := config.GetConfig()

	log.Print("App successfully started")

	app, err := app.NewApp(config)

    if err != nil {
        
    }
}
