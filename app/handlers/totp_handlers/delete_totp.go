package totp_handlers

import (
	"authentication_backend/app/actions/totp_actions"
	"authentication_backend/app/models/user_models"
	"authentication_backend/utils/log"
	"authentication_backend/utils/response"
	"net/http"
)

func DeleteTOTP(w http.ResponseWriter, r *http.Request) {
	log.Api(r)

	userID := r.PathValue("userId")

	user := user_models.GetUserByID(userID)
	if user == nil {
		response.NewErrorMessage(w, response.ErrUserNotFound, http.StatusInternalServerError)
		return
	}

	if err := totp_actions.DisableTOTP(user); err != nil {
		response.NewErrorMessage(w, response.ErrDisableTOTP, http.StatusInternalServerError)
		return
	}

	response.NewSuccessMessage(w, response.SuccessDisableTOTP)
}
