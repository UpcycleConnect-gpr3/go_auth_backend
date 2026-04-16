package user_actions

import (
	"authentication_backend/app/models/totp_models"
	"authentication_backend/app/models/user_models"
	"authentication_backend/utils/jwt"
	"authentication_backend/utils/response"
	"authentication_backend/utils/rules"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

func loginValidateCredential(userDto user_models.Credentials) ([]rules.ValidationError, *user_models.User) {
	var errs []rules.ValidationError

	rules.StringMinLength(userDto.Email, 5, "email", &errs)
	rules.StringMinLength(userDto.Password, 6, "password", &errs)
	rules.StringMaxLength(userDto.Password, 30, "password", &errs)

	existing := user_models.GetUserBy([]string{"id", "email", "password", "totp_enabled"}, "email = ?", userDto.Email)
	if existing == nil {
		errs = append(errs, rules.ValidationError{
			Field:   "email",
			Message: response.ErrAuthFailed,
		})
		return errs, existing
	}

	isCorrectPassword := existing.CheckPassword(userDto.Password)

	if !isCorrectPassword || userDto.Email != existing.Email {
		errs = append(errs, rules.ValidationError{
			Field:   "email",
			Message: response.ErrAuthFailed,
		})
	}

	return errs, existing
}

func Login(credentials user_models.Credentials) (string, bool, []rules.ValidationError, *user_models.User) {

	validationErrors, existing := loginValidateCredential(credentials)
	if len(validationErrors) > 0 {
		return "", false, validationErrors, existing
	}

	if existing.TOTPEnabled {

		timestamp := time.Now().Unix()
		hashInput := fmt.Sprintf("%s:%d:%s", existing.Id.String(), timestamp, jwt.PrivateKey)
		hash := sha256.Sum256([]byte(hashInput))
		hashStr := hex.EncodeToString(hash[:])

		if err := totp_models.CreateTOTPHash(existing.Id.String(), hashStr); err != nil {
			return "", false, []rules.ValidationError{
				{Field: "totp", Message: response.ErrGenerateTOTP},
			}, existing
		}

		return hashStr, true, nil, existing
	}

	return "", false, nil, existing
}
