package user_actions

import (
	"authentication_backend/app/models/user_models"
	"authentication_backend/utils/rules"
)

func createValidateUser(userDto user_models.Credentials) []rules.ValidationError {
	var errs []rules.ValidationError
	var user user_models.User

	rules.StringMinLength(userDto.Email, 5, "email", &errs)
	rules.StringMinLength(userDto.Password, 6, "password", &errs)
	rules.StringMaxLength(userDto.Password, 128, "password", &errs)
	rules.MustContainsAny(userDto.Password, "!@#$%^&*()", 1, "password", &errs)

	err := user.Get([]string{"id", "email"}, "email = ?", userDto.Email)
	if err != nil {
		errs = append(errs, rules.ValidationError{
			Field:   "email",
			Message: "email must be unique",
		})
	}

	return errs
}

func CreateUser(userDto user_models.Credentials) (*user_models.User, []rules.ValidationError) {

	validationError := createValidateUser(userDto)

	if len(validationError) > 0 {
		return nil, validationError
	}

	user := user_models.CreateUser(userDto)

	return user, nil
}
