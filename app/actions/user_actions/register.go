package user_actions

import (
	"authentication_backend/app/models/user_models"
	"authentication_backend/utils/rules"
)

func createValidateUser(userDto user_models.Credentials) []rules.ValidationError {
	var errs []rules.ValidationError

	rules.StringMinLength(userDto.Email, 5, "email", &errs)
	rules.StringMinLength(userDto.Password, 6, "password", &errs)
	rules.StringMaxLength(userDto.Password, 30, "password", &errs)
	rules.MustContainsAny(userDto.Password, "!@#$%^&*()", 1, "password", &errs)

	existing := user_models.GetUserByEmailAuth(userDto.Email)
	if existing != nil {
		errs = append(errs, rules.ValidationError{
			Field:   "email",
			Message: "email must be unique",
		})
	}

	return errs
}

func CreateUser(userDto user_models.Credentials) []rules.ValidationError {

	validationError := createValidateUser(userDto)

	if len(validationError) > 0 {
		return validationError
	}

	user_models.CreateUser(userDto)

	return nil
}
