package middlewares

import (
	"crypto/tls"
	"time"

	"github.com/thirumathikart/thirumathikart-messaging-service/config"
	mail "github.com/xhit/go-simple-mail/v2"
)

func SetupMailer() *mail.SMTPServer {

	mailer := mail.NewSMTPClient()
	mailer.Host = config.MailHost
	mailer.Port = config.MailPort
	mailer.Username = config.MailUser
	mailer.Password = config.MailPassword
	mailer.Encryption = mail.EncryptionSTARTTLS
	mailer.KeepAlive = false
	mailer.ConnectTimeout = 10 * time.Second
	mailer.SendTimeout = 10 * time.Second
	mailer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	return mailer

}
