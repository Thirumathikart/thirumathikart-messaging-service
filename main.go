package main

import (
	"fmt"
	"log"
	"net"

	"github.com/thirumathikart/thirumathikart-messaging-service/config"
	"github.com/thirumathikart/thirumathikart-messaging-service/controllers"
	"github.com/thirumathikart/thirumathikart-messaging-service/middlewares"
	mail "github.com/thirumathikart/thirumathikart-messaging-service/rpcs/mail"
	notification "github.com/thirumathikart/thirumathikart-messaging-service/rpcs/notification"
	"google.golang.org/grpc"
)

func main() {
	url := config.InitConfig()

	lis, err := net.Listen("tcp", *url)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(middlewares.WithServerUnaryInterceptor())

	notification.RegisterNotificationServiceServer(grpcServer, &controllers.NotificationRPCServer{})
	mail.RegisterMailServiceServer(grpcServer, &controllers.MailRPCServer{})

	fmt.Println("starting grpc server on", *url)

	err1 := grpcServer.Serve(lis)
	if err1 != nil {
		log.Panic("grpc server running error on", err1)
	}
}
