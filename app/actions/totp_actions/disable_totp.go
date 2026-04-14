package totp_actions

import (
	"authentication_backend/app/models/user_models"
	"authentication_backend/utils/log"
)

func DisableTOTP(user *user_models.User) error {
	user.TOTPEnabled = false
	user.TOTPSecret = ""
	err := user_models.UpdateUserTOTP(user)
	if err != nil {
		log.Database("DisableTOTP", err)
		return err
	}
	return nil
}
