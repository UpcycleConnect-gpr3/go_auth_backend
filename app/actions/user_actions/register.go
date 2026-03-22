package user_actions

import (
	"authentication_backend/app/models/user_models"
	"authentication_backend/utils/log"
	"authentication_backend/utils/rules"
	"encoding/json"
	"net/http"
)

func createValidateUser(userDto user_models.Credentials) []string {
	var errsMsg []string

	rules.StringMinLength(userDto.Email, 5, "email", &errsMsg)
	rules.StringMinLength(userDto.Password, 6, "password", &errsMsg)

	/*existing, err := models.GetUserByEmail(userDto.Email)

	if err != nil {
		errsMsg = append(errsMsg, "Error getting User by username")
	}

	if existing != nil {
		errsMsg = append(errsMsg, "Username must be unique")
	}*/

	return errsMsg
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var userDto user_models.Credentials

	err := json.NewDecoder(r.Body).Decode(&userDto)
	log.Info(userDto.Email)
	log.Info(userDto.Password)
	if err != nil {
		http.Error(w, "Incorrect body format", http.StatusBadRequest)
		return
	}

	errMsg := createValidateUser(userDto)

	if len(errMsg) > 0 {
		encoded, _ := json.Marshal(errMsg)
		http.Error(w, string(encoded), http.StatusBadRequest)
		return
	}

	log.Api(r)

	err = user_models.CreateUser(userDto)

	if err != nil {
		log.Info(err.Error())
		http.Error(w, "pb insertion", http.StatusInternalServerError)
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
