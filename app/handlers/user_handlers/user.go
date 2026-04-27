package user_handlers

import (
	"authentication_backend/app/models/user_models"
	"authentication_backend/utils/log"
	"authentication_backend/utils/request"
	"authentication_backend/utils/response"
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
	// TODO: Implement
	response.NewSuccessMessage(w, "User updated successfully")
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Api(r)
	// TODO: Implement
	response.NewSuccessMessage(w, "User deleted successfully")
}
