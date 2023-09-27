package main

import (
	"context"
	"firebase.google.com/go/v4/messaging"
	"github.com/joho/godotenv"
	"github.com/qcodelabsllc/qreeket/notification/config"
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

	// test the messaging client
	sendResult, err := config.FirebaseMessaging.Send(context.Background(), &messaging.Message{
		Notification: &messaging.Notification{
			Title: "Test",
			Body:  "This is a test notification",
		},
		Topic: "test",
	})
	if err != nil {
		log.Fatalf("unable to send message: %+v\n", err)
	}

	log.Printf("message sent: %+v\n", sendResult)
}
