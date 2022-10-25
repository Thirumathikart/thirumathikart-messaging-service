package config

import (
	"time"

	mail "github.com/xhit/go-simple-mail/v2"
)

var smtpClient *mail.SMTPClient

func SetupMailer() {

	mailer := mail.NewSMTPClient()
	mailer.Host = MailHost
	mailer.Port = MailPort
	mailer.Username = MailUser
	mailer.Password = MailPassword
	mailer.Encryption = mail.EncryptionSSL
	mailer.KeepAlive = false
	mailer.Authentication = mail.AuthPlain
	mailer.ConnectTimeout = 10 * time.Second
	mailer.SendTimeout = 10 * time.Second
	var err error

	smtpClient, err = mailer.Connect()
	if err != nil {
		GrpcLog.Error(err)
	}
}

func GetMailer() mail.SMTPClient {
	return *smtpClient
}
