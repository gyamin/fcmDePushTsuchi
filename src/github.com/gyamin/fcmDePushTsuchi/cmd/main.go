package main

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
)

func main() {
	// PUSHテスト
	var app *firebase.App = initializeServiceAccountID()
	sendToToken(app)
}

func initializeServiceAccountID() *firebase.App {
	// [START initialize_sdk_with_service_account_id]
	conf := &firebase.Config{
		ServiceAccountID: "my-client-id@my-project-id.iam.gserviceaccount.com",
	}
	app, err := firebase.NewApp(context.Background(), conf)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	// [END initialize_sdk_with_service_account_id]
	return app
}

func sendToToken(app *firebase.App) {
	// [START send_to_token_golang]
	// Obtain a messaging.Client from the App.
	ctx := context.Background()
	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}

	// This registration token comes from the client FCM SDKs.
	registrationToken := "cS62HS9Mmtuff7tv1irQJQ:APA91bGekIxZwT0Z_WHwnx7Broef8qPweK2qhW9r5ZMktd3iyZpoUQNxBrsrC68wgmUsGVMhyLsatoG01Nh_wlXpAxWi-vnCGB0FbkpycrneNZmYRdyqexKllR9o98duJDb7juoOL-Im"

	// See documentation on defining a message payload.
	message := &messaging.Message{
		Data: map[string]string{
			"score": "850",
			"time":  "2:45",
		},
		Token: registrationToken,
	}

	// Send a message to the device corresponding to the provided
	// registration token.
	response, err := client.Send(ctx, message)
	if err != nil {
		log.Fatalln(err)
	}
	// Response is a message ID string.
	fmt.Println("Successfully sent message:", response)
	// [END send_to_token_golang]
}
