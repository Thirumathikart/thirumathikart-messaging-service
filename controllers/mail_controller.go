package controllers

import (
	"context"
	"log"

	"github.com/thirumathikart/thirumathikart-messaging-service/config"
	mail "github.com/thirumathikart/thirumathikart-messaging-service/rpcs/mail"
	simplemail "github.com/xhit/go-simple-mail/v2"
)

type MailRPCServer struct {
	mail.UnimplementedMailServiceServer
}

func (MailRPCServer) SendSingleMailRPC(ctx context.Context, request *mail.SingleMailRequest) (*mail.SingleMailResponse, error) {

	smtpClient := config.GetMailer()

	email := simplemail.NewMSG().SetFrom("From Thirumathikart <thirumathikart@nitt.edu>").
		AddTo(request.Reciever).
		SetSubject(request.Subject).SetBody(simplemail.TextPlain, request.Body)

	if request.File != nil {
		for _, file := range request.File {
			email.Attach(&simplemail.File{Data: file.File, Name: file.Name})
		}
	}
	// always check error after send
	if email.Error != nil {
		log.Fatal(email.Error)
	}

	err := email.Send(&smtpClient)
	IsSuccess := true
	if err != nil {
		IsSuccess = false
	}

	return &mail.SingleMailResponse{IsSuccess: IsSuccess}, err

}
