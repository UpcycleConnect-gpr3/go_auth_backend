package totp_actions

import (
	"authentication_backend/app/models/user_models"
	"authentication_backend/utils/log"

	"github.com/pquerna/otp/totp"
)

func GenerateAndStoreTOTP(user *user_models.User) (string, error) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "Upcycle Connect",
		AccountName: user.Email,
	})
	if err != nil {
		log.Database("GenerateAndStoreTOTP", err)
		return "", err
	}
	user.TOTPEnabled = false
	user.TOTPSecret = key.Secret()

	if err := user_models.UpdateUserTOTP(user); err != nil {
		log.Database("UpdateUserTOTP (store secret)", err)
		return "", err
	}
	return key.URL(), nil
}
