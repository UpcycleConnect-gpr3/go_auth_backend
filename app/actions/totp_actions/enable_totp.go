package totp_actions

import (
	"authentication_backend/app/models/user_models"
	"authentication_backend/utils/log"
)

func EnableTOTP(user *user_models.User) error {
	user.TOTPEnabled = true
	if err := user_models.UpdateUserTOTP(user); err != nil {
		log.Database("EnableTOTP", err)
		return err
	}
	return nil
}
