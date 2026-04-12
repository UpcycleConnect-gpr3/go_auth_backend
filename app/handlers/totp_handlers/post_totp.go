package totp_handlers

import (
	"authentication_backend/app/actions/totp_actions"
	"authentication_backend/app/models/totp_models"
	"authentication_backend/app/models/user_models"
	"authentication_backend/utils/log"
	"authentication_backend/utils/response"
	"encoding/json"
	"net/http"
)

func PostTOTP(w http.ResponseWriter, r *http.Request) {
	log.Api(r)

	userID := r.PathValue("userId")

	request := totp_models.TOTPCodeRequest{}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		response.NewErrorMessage(w, response.ErrJson, http.StatusBadRequest)
		return
	}

	user := user_models.GetUserByIDWithTOTP(userID)
	if user == nil {
		response.NewErrorMessage(w, response.ErrUserNotFound, http.StatusNotFound)
		return
	}

	if !totp_actions.ValidateTOTP(user, request.Code) {
		response.NewErrorMessage(w, response.ErrInvalidTOTP, http.StatusBadRequest)
		return
	}

	if err := totp_actions.EnableTOTP(user); err != nil {
		response.NewErrorMessage(w, response.ErrEnableTOTP, http.StatusInternalServerError)
		return
	}

	response.NewSuccessMessage(w, response.SuccessEnableTOTP)
}
