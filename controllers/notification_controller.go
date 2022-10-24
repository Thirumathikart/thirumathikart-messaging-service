package controllers

import (
	"context"

	firebase "firebase.google.com/go/v4"
	messaging "firebase.google.com/go/v4/messaging"
	"github.com/thirumathikart/thirumathikart-messaging-service/middlewares"
	notification "github.com/thirumathikart/thirumathikart-messaging-service/rpcs/notification"
	"google.golang.org/api/option"
)

type NotificationRPCServer struct {
	notification.UnimplementedNotificationServiceServer
}

func (NotificationRPCServer) SendSinglePushNotification(ctx context.Context, request *notification.SingleNotificationRequest) (*notification.SingleNotificationResponse, error) {
	decodedKey, err := middlewares.DecodedFireBaseKey()

	if err != nil {
		return &notification.SingleNotificationResponse{IsSuccess: false}, err
	}

	opts := []option.ClientOption{option.WithCredentialsJSON(decodedKey)}

	app, err := firebase.NewApp(context.Background(), nil, opts...)

	if err != nil {
		return &notification.SingleNotificationResponse{IsSuccess: false}, err
	}

	fcmClient, err := app.Messaging(context.Background())

	if err != nil {
		return &notification.SingleNotificationResponse{IsSuccess: false}, err
	}

	_, err = fcmClient.Send(context.Background(), &messaging.Message{

		Notification: &messaging.Notification{
			Title: request.Title,
			Body:  request.Body,
		},
		Token: request.Token, // it's a single device token
	})

	if err != nil {
		return &notification.SingleNotificationResponse{IsSuccess: false}, err
	}

	return &notification.SingleNotificationResponse{IsSuccess: true}, nil
}

func (NotificationRPCServer) SendMultiplePushNotification(ctx context.Context, request *notification.MultipleNotificationRequest) (*notification.MultipleNotificationResponse, error) {
	decodedKey, err := middlewares.DecodedFireBaseKey()

	if err != nil {
		return nil, err
	}

	opts := []option.ClientOption{option.WithCredentialsJSON(decodedKey)}

	app, err := firebase.NewApp(context.Background(), nil, opts...)

	if err != nil {
		return nil, err
	}

	fcmClient, err := app.Messaging(context.Background())

	if err != nil {
		return nil, err
	}

	response, err := fcmClient.SendMulticast(context.Background(), &messaging.MulticastMessage{
		Notification: &messaging.Notification{
			Title: request.Title,
			Body:  request.Body,
		},
		Tokens: request.Token,
	})

	if err != nil {
		return nil, err
	}
	return &notification.MultipleNotificationResponse{
		SuccessCount: int64(response.SuccessCount),
		FailureCount: int64(response.FailureCount),
	}, nil
}
