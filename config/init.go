package config

import (
	"flag"

	"github.com/thirumathikart/thirumathikart-messaging-service/models"
)

func InitConfig() *string {
	SetupLogger()
	SetupEnvironment()
	SetupFCM()
	SetupMailer()

	var flags *models.Flags
	flag.Parse()
	port := flag.String("port", ServerPort, "port to be used")
	ip := flag.String("ip", ServerHost, "ip to be used")

	flags = models.NewFlags(*ip, *port)
	url, _ := flags.GetApplicationURL()

	return url
}
