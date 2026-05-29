package user_actions

import (
	"authentication_backend/app/models/user_models"
)

func DeleteUser(userId string) error {
	var user user_models.User
	err := user.Get([]string{"id"}, "id = ?", userId)
	if err != nil {
		return err
	}

	err = user.Delete()
	if err != nil {
		return err
	}

	return nil
}
