package user_actions

import (
	"authentication_backend/app/models/user_models"
	"authentication_backend/utils/db"
	"authentication_backend/utils/rules"
	"authentication_backend/var/database"
)

func validateUpdateUser(userDto user_models.UpdateUserRequest, userId string) []rules.ValidationError {
	var errs []rules.ValidationError

	if userDto.Email != nil {
		rules.StringMinLength(*userDto.Email, 5, "email", &errs)
		rules.Unique[user_models.User](database.Auth, user_models.TABLE, "email", "email = ? AND id <> ?", &errs, userDto.Email, userId)
	}

	if userDto.Password != nil {
		rules.StringMinLength(*userDto.Password, 6, "password", &errs)
		rules.StringMaxLength(*userDto.Password, 128, "password", &errs)
		rules.MustContainsAny(*userDto.Password, "!@#$%^&*()", 1, "password", &errs)
	}

	return errs
}

func UpdateUser(userId string, userDto user_models.UpdateUserRequest) (*user_models.User, []rules.ValidationError) {
	validationErrors := validateUpdateUser(userDto, userId)
	if len(validationErrors) > 0 {
		return nil, validationErrors
	}

	var user user_models.User

	err := user.Update(userDto, userId)
	if err != nil {
		return nil, []rules.ValidationError{
			{Field: "email", Message: "Failed to update user"},
		}
	}

	err = user.Get([]string{"id", "email", "firstname", "lastname", "totp_enabled"}, db.IdClause, userId)
	if err != nil {
		return nil, []rules.ValidationError{
			{Field: "email", Message: "Failed to get user"},
		}
	}

	return &user, nil
}
