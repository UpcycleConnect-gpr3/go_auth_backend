package totp_handlers

import (
	"authentication_backend/app/actions/totp_actions"
	"authentication_backend/utils/auth"
	"authentication_backend/utils/log"
	"authentication_backend/utils/response"
	"net/http"
)

func GetTOTP(w http.ResponseWriter, r *http.Request) {
	log.Api(r)

	user := auth.Auth(r).User(w, []string{"id", "email"})

	totpURL, err := totp_actions.GenerateAndStoreTOTP(user)
	if err != nil {
		response.NewErrorMessage(w, response.ErrGenerateTOTP, http.StatusInternalServerError)
		return
	}

	response.NewSuccessData(w, map[string]string{"totp_url": totpURL})
}
