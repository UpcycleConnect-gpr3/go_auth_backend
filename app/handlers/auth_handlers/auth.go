package auth_handlers

import (
	"authentication_backend/app/actions/user_actions"
	"authentication_backend/app/models/user_models"
	"authentication_backend/utils/jwt"
	"encoding/json"
	"fmt"
	"net/http"
)

type TokenResponse struct {
	BearerToken string `json:"bearer_token"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var creds user_models.Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	existing, err := user_models.GetUserByEmail(creds.Email)

	if err != nil {
		http.Error(w, "Error requesting user_actions by username", http.StatusInternalServerError)
		return
	}

	isCorrectPassword := existing.CheckPassword(creds.Password)

	if !isCorrectPassword || creds.Email != existing.Email {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	token, err := jwt.GenerateJWT(existing.Id.String())
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	tokenResponse := TokenResponse{BearerToken: token}
	encodedToken, _ := json.Marshal(tokenResponse)
	fmt.Fprintf(w, "%s", encodedToken)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// TODO : Write Logout handler with user_actions
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	user_actions.CreateUser(w, r)
}
