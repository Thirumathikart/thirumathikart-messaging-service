package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/thirumathikart/thirumathikart-messaging-service/config"
	"github.com/thirumathikart/thirumathikart-messaging-service/controllers"
	"github.com/thirumathikart/thirumathikart-messaging-service/models"
	mail "github.com/thirumathikart/thirumathikart-messaging-service/rpcs/mail"
	notification "github.com/thirumathikart/thirumathikart-messaging-service/rpcs/notification"
	"google.golang.org/grpc"
)

func main() {
	var flags *models.Flags
	flag.Parse()

	port := flag.String("port", config.ServerPort, "port to be used")
	ip := flag.String("ip", config.ServerHost, "ip to be used")

	flags = models.NewFlags(*ip, *port)
	url, _ := flags.GetApplicationURL()

	lis, err := net.Listen("tcp", *url)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)

	notification.RegisterNotificationServiceServer(grpcServer, &controllers.NotificationRPCServer{})
	mail.RegisterMailServiceServer(grpcServer, &controllers.MailRPCServer{})

	fmt.Println("starting grpc server on", *url)

	err1 := grpcServer.Serve(lis)
	if err1 != nil {
		fmt.Println("grpc server running error on", err1)
	}
}
