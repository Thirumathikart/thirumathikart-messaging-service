package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/fatih/color"

	"github.com/joho/godotenv"
)

var ServerPort string
var ServerHost string
var MailPort int
var MailHost string
var MailUser string
var MailPassword string
var FireBaseAuthKey string

func InitConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(color.RedString("Error loading .env"))
	}

	ServerPort = os.Getenv("SERVER_PORT")
	ServerHost = os.Getenv("SERVER_HOST")
	MailHost = os.Getenv("MAIL_HOST")
	MailUser = os.Getenv("MAIL_USER")
	MailPassword = os.Getenv("MAIL_PASSWORD")
	FireBaseAuthKey = os.Getenv("FIREBASE_AUTH_KEY")
	MailPort, err = strconv.Atoi(os.Getenv("MAIL_PORT"))
	if err != nil {
		// ... handle error
		panic(err)
	}
}
