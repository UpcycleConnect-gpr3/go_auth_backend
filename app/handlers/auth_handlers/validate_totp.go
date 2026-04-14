package auth_handlers

import (
	"authentication_backend/app/actions/user_actions"
	"authentication_backend/utils/log"
	"authentication_backend/utils/response"
	"encoding/json"
	"net/http"
)

func ValidateTOTPHandler(w http.ResponseWriter, r *http.Request) {
	log.Api(r)

	var request struct {
		Hash string `json:"hash"`
		Code string `json:"code"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		response.NewErrorMessage(w, response.ErrJson, http.StatusBadRequest)
		return
	}

	token, err := user_actions.ValidateTOTPAndLogin(request.Hash, request.Code)
	if err != nil {
		response.NewErrorMessage(w, err.Error(), http.StatusUnauthorized)
		return
	}

	response.NewSuccessData(w, map[string]string{
		"bearer_token": token,
	})
}
