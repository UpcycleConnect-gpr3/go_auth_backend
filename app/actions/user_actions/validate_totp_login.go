package user_actions

import (
	"authentication_backend/app/actions/totp_actions"
	"authentication_backend/app/models/totp_models"
	"authentication_backend/app/models/user_models"
	"authentication_backend/utils/jwt"
	"authentication_backend/utils/response"
	"fmt"
)

func ValidateTOTPAndLogin(hash, totpCode string) (string, error) {
	totpRecord, err := totp_models.GetTOTPByHash(hash)
	if err != nil {
		return "", fmt.Errorf(response.ErrFetchingTOTPRecord)
	}
	if totpRecord == nil {
		return "", fmt.Errorf(response.ErrInvalidOrExpiredHash)
	}

	user := user_models.GetUserByID(totpRecord.UserID)
	if user == nil {
		return "", fmt.Errorf(response.ErrUserNotFound)
	}

	if !totp_actions.ValidateTOTP(user, totpCode) {
		return "", fmt.Errorf(response.ErrInvalidTOTP)
	}

	if err = totp_models.DeleteTOTPByHash(hash); err != nil {
		return "", err
	}

	token, err := jwt.GenerateJWT(totpRecord.UserID)
	if err != nil {
		return "", fmt.Errorf(response.ErrGenerateToken)
	}

	return token, nil
}
