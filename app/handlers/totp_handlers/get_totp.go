package totp_handlers

import (
	"authentication_backend/app/actions/totp_actions"
	"authentication_backend/app/models/user_models"
	"authentication_backend/utils/log"
	"authentication_backend/utils/response"
	"net/http"
)

func GetTOTP(w http.ResponseWriter, r *http.Request) {
	log.Api(r)

	userID := r.PathValue("userId")

	user := user_models.GetUserByID(userID)
	if user == nil {
		response.NewErrorMessage(w, response.ErrUserNotFound, http.StatusNotFound)
		return
	}

	totpURL, err := totp_actions.GenerateAndStoreTOTP(user)
	if err != nil {
		response.NewErrorMessage(w, response.ErrGenerateTOTP, http.StatusInternalServerError)
		return
	}

	response.NewSuccessData(w, map[string]string{"totp_url": totpURL})
}
