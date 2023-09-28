package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/qcodelabsllc/qreeket/notification/config"
	"github.com/qcodelabsllc/qreeket/notification/network"
	"log"
)

// entry point of the application
func main() {
	// This line loads the environment variables from the ".env" file.
	if err := godotenv.Load(); err != nil {
		log.Fatalf("unable to load environment variables: %+v\n", err)
	}

	// initialize the messaging client
	if err := config.InitFirebase(context.Background()); err != nil {
		log.Fatalf("unable to initialize firebase: %+v\n", err)
	}

	// initialize the grpc server
	network.InitServer()
}
