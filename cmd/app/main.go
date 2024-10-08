package main

import (
	"log"
	"server/internal/app"
)

func main() {
	application, err := app.NewApp()
	if err != nil {
		log.Fatalf("failed to init app: %s", err.Error())
	}
	err = application.Run()
	if err != nil {
		log.Fatalf("failed to run app: %s", err.Error())
	}
}
