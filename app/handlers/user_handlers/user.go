package user_handlers

import (
	"authentication_backend/app/actions/user_actions"
	"authentication_backend/app/models/user_models"
	"authentication_backend/utils/auth"
	"authentication_backend/utils/log"
	"authentication_backend/utils/request"
	"authentication_backend/utils/response"
	"encoding/json"
	"net/http"
)

func IndexUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Api(r)

	var users []user_models.User
	var user *user_models.User
	columns := []string{"id", "email", "firstname", "lastname", "created_at", "updated_at"}

	err := user.All(columns, &users)
	if err != nil {
		response.NewErrorMessage(w, response.ErrUserNotFound, http.StatusInternalServerError)
		return
	}

	response.NewSuccessData(w, users)
}

func ShowUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Api(r)

	id := request.Request(r, "id").Value()

	var user user_models.User
	columns := []string{"id", "firstname", "lastname", "email", "created_at", "updated_at"}

	err := user.Get(columns, "id = ?", id)

	if err != nil {
		response.NewErrorMessage(w, response.ErrUserNotFound, http.StatusNotFound)
		return
	}

	response.NewSuccessData(w, user)
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Api(r)

	userId := auth.Auth(r).Id()

	var updateRequest user_models.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&updateRequest); err != nil {
		response.NewErrorMessage(w, response.ErrJson, http.StatusBadRequest)
		return
	}

	user, validationErrors := user_actions.UpdateUser(userId, updateRequest)
	if len(validationErrors) > 0 {
		response.NewValidationError(w, response.ErrInvalidBody, validationErrors)
		return
	}

	response.NewSuccessData(w, user)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Api(r)

	userId := auth.Auth(r).Id()

	if err := user_actions.DeleteUser(userId); err != nil {
		response.NewErrorMessage(w, response.ErrUserNotFound, http.StatusNotFound)
		return
	}

	response.NewSuccessMessage(w, response.SuccessUserDeleted)
}
