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

type TokenResponse struct {
	BearerToken string `json:"bearer_token"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	log.Api(r)

	var credentials user_models.Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)

	if err != nil {
		response.NewErrorMessage(w, response.ErrJson, http.StatusBadRequest)
		return
	}

	validationErrors, existing := user_actions.Login(credentials)

	if len(validationErrors) > 0 {
		response.NewValidationError(w, response.ErrInvalidBody, validationErrors)
		return
	}

	token, err := jwt.GenerateJWT(existing.Id.String())
	if err != nil {
		response.NewErrorMessage(w, "Échec de la génération du token", http.StatusInternalServerError)
		return
	}

	response.NewSuccessData(w, TokenResponse{BearerToken: token}, response.SuccessLogin)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	log.Api(r)

	w.Header().Set("Content-Type", "application/json")

	var userDto user_models.Credentials

	err := json.NewDecoder(r.Body).Decode(&userDto)

	if err != nil {
		response.NewErrorMessage(w, response.ErrInvalidBody, http.StatusBadRequest)
	}

	validationErrors := user_actions.CreateUser(userDto)

	if len(validationErrors) > 0 {
		response.NewValidationError(w, response.ErrInvalidBody, validationErrors)
		return
	}

	w.WriteHeader(http.StatusCreated)

	/*	token, err := jwt.GenerateJWT(existing.Id.String())
		if err != nil {
			http.Error(w, "Error generating token", http.StatusInternalServerError)
			return
		}

		tokenResponse := TokenResponse{BearerToken: token}
		encodedToken, _ := json.Marshal(tokenResponse)
		fmt.Fprintf(w, "%s", encodedToken)*/
}
