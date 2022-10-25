package config

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"google.golang.org/api/option"
)

var fireClient *messaging.Client

func SetupFCM() {

	fireBaseAuthKey := FireBaseAuthKey

	opts := []option.ClientOption{option.WithCredentialsJSON([]byte(fireBaseAuthKey))}

	app, err := firebase.NewApp(context.Background(), nil, opts...)
	if err != nil {
		panic(err)
	}
	fireClient, err = app.Messaging(context.Background())

	if err != nil {
		panic(err)
	}

}

func GetFCM() (*messaging.Client, error) {
	var err error
	if err != nil {
		fmt.Print(err)
	}
	return fireClient, err
}
