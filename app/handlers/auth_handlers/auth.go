package auth_handlers

import (
	"authentication_backend/app/actions/user_actions"
	"authentication_backend/app/models/user_models"
	"authentication_backend/utils/jwt"
	"authentication_backend/utils/log"
	"authentication_backend/utils/response"
	"encoding/json"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	log.Api(r)

	var credentials user_models.Credentials
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		response.NewErrorMessage(w, response.ErrJson, http.StatusBadRequest)
		return
	}

	hash, totpRequired, validationErrors, user := user_actions.Login(credentials)

	if len(validationErrors) > 0 {
		response.NewValidationError(w, response.ErrInvalidBody, validationErrors)
		return
	}

	if totpRequired {
		response.NewSuccessData(w, map[string]interface{}{
			"hash":          hash,
			"totp_required": true,
		})
		return
	}

	token, err := jwt.GenerateJWT(user.Id.String())
	if err != nil {
		response.NewErrorMessage(w, response.ErrGenerateToken, http.StatusInternalServerError)
		return
	}

	response.NewSuccessData(w, map[string]interface{}{
		"bearer_token":  token,
		"totp_required": false,
	})
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	log.Api(r)

	w.Header().Set("Content-Type", "application/json")

	var userDto user_models.Credentials

	err := json.NewDecoder(r.Body).Decode(&userDto)

	if err != nil {
		response.NewErrorMessage(w, response.ErrInvalidBody, http.StatusBadRequest)
	}

	user, validationErrors := user_actions.CreateUser(userDto)

	if len(validationErrors) > 0 {
		response.NewValidationError(w, response.ErrInvalidBody, validationErrors)
		return
	}

	w.WriteHeader(http.StatusCreated)
	response.NewSuccessData(w, map[string]interface{}{
		"user_id": user.Id,
	})
}
