package controllers

import (
	"context"

	messaging "firebase.google.com/go/v4/messaging"
	"github.com/thirumathikart/thirumathikart-messaging-service/config"
	notification "github.com/thirumathikart/thirumathikart-messaging-service/rpcs/notification"
)

type NotificationRPCServer struct {
	notification.UnimplementedNotificationServiceServer
}

func (NotificationRPCServer) SendSingleNotificationRPC(ctx context.Context, request *notification.SingleNotificationRequest) (*notification.SingleNotificationResponse, error) {
	fcmClient, err := config.GetFCM()

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

func (NotificationRPCServer) SendMultipleNotificationRPC(ctx context.Context, request *notification.MultipleNotificationRequest) (*notification.MultipleNotificationResponse, error) {
	fcmClient, err := config.GetFCM()

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
