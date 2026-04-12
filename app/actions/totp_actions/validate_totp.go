package totp_actions

import (
	"authentication_backend/app/models/user_models"

	"github.com/pquerna/otp/totp"
)

func ValidateTOTP(user *user_models.User, code string) bool {
	return totp.Validate(code, user.TOTPSecret)
}
