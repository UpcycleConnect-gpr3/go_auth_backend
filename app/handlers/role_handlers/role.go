package role_handlers

import (
	"authentication_backend/app/actions/role_actions"
	"authentication_backend/app/models/user_models"
	"authentication_backend/utils/auth"
	"authentication_backend/utils/log"
	"authentication_backend/utils/request"
	"authentication_backend/utils/response"
	"encoding/json"
	"net/http"
)

func UpdateUserRoleHandler(w http.ResponseWriter, r *http.Request) {
	log.Api(r)

	adminId := auth.Auth(r).Id()
	targetUserId := request.Request(r, "id").Value()
	if targetUserId == "" {
		response.NewErrorMessage(w, response.ErrUserNotFound, http.StatusNotFound)
		return
	}

	var roleUpdate struct {
		Role user_models.Role `json:"role"`
	}
	if err := json.NewDecoder(r.Body).Decode(&roleUpdate); err != nil {
		response.NewErrorMessage(w, response.ErrJson, http.StatusBadRequest)
		return
	}

	validationErrors := role_actions.UpdateUserRole(adminId, targetUserId, roleUpdate.Role)
	if len(validationErrors) > 0 {
		response.NewValidationError(w, response.ErrInvalidBody, validationErrors)
		return
	}

	response.NewSuccessMessage(w, "Role updated successfully")
}

func GetAllRolesHandler(w http.ResponseWriter, r *http.Request) {
	log.Api(r)

	response.NewSuccessData(w, map[string]interface{}{
		"roles": user_models.AllRoles(),
	})
}

func GetUserSelectableRolesHandler(w http.ResponseWriter, r *http.Request) {
	log.Api(r)

	response.NewSuccessData(w, map[string]interface{}{
		"user_selectable_roles": user_models.AllUserSelectableRoles(),
	})
}
