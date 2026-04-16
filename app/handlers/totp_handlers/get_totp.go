package totp_handlers

import (
	"authentication_backend/app/actions/totp_actions"
	"authentication_backend/app/middleware/auth_middleware"
	"authentication_backend/app/models/user_models"
	"authentication_backend/utils/log"
	"authentication_backend/utils/response"
	"net/http"
)

func GetTOTP(w http.ResponseWriter, r *http.Request) {
	log.Api(r)

	userID := auth_middleware.GetUserId(r.Context())

	user := user_models.GetUserBy([]string{"id", "email"}, "id = ?", userID)
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
