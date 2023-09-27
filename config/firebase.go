package config

import (
	"context"
	"errors"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"google.golang.org/api/option"
	"path/filepath"
	"time"
)

var (
	// FirebaseMessaging is the firebase messaging client
	FirebaseMessaging *messaging.Client
)

// InitFirebase initializes the firebase app and messaging client
func InitFirebase(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// load service key file from absolute path
	serviceAccountKey := "./qreeket-firebase-service-key.json"
	serviceAccountKeyPath, err := filepath.Abs(serviceAccountKey)
	if err != nil {
		return errors.New("failed to load service account key file")
	}

	// initialize firebase app
	app, err := firebase.NewApp(ctx, nil, option.WithCredentialsFile(serviceAccountKeyPath))
	if err != nil {
		return errors.New("failed to initialize firebase app")
	}

	// initialize messaging client
	msg, err := app.Messaging(ctx)

	// set global variables
	FirebaseMessaging = msg

	return nil
}
