package middlewares

import (
	"encoding/base64"

	"github.com/thirumathikart/thirumathikart-messaging-service/config"
)

func DecodedFireBaseKey() ([]byte, error) {

	fireBaseAuthKey := config.FireBaseAuthKey

	decodedKey, err := base64.StdEncoding.DecodeString(fireBaseAuthKey)
	if err != nil {
		return nil, err
	}
	return decodedKey, err
}
