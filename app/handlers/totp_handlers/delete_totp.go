package totp_handlers

import (
	"authentication_backend/app/actions/totp_actions"
	"authentication_backend/utils/auth"
	"authentication_backend/utils/log"
	"authentication_backend/utils/response"
	"net/http"
)

func DeleteTOTP(w http.ResponseWriter, r *http.Request) {
	log.Api(r)

	user := auth.Auth(r).User(w, []string{})

	if err := totp_actions.DisableTOTP(user); err != nil {
		response.NewErrorMessage(w, response.ErrDisableTOTP, http.StatusInternalServerError)
		return
	}

	response.NewSuccessMessage(w, response.SuccessDisableTOTP)
}
