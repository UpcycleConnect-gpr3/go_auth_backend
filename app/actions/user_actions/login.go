package user_actions

import (
	"authentication_backend/app/models/user_models"
	"authentication_backend/utils/log"
	"authentication_backend/utils/rules"
)

func loginValidateCredential(userDto user_models.Credentials) ([]rules.ValidationError, *user_models.User) {
	var errs []rules.ValidationError

	rules.StringMinLength(userDto.Email, 5, "email", &errs)
	rules.StringMinLength(userDto.Password, 6, "password", &errs)
	rules.StringMaxLength(userDto.Password, 30, "password", &errs)

	existing := user_models.GetUserByEmail(userDto.Email)
	if existing == nil {
		errs = append(errs, rules.ValidationError{
			Field:   "email",
			Message: log.ErrAuthFailed,
		})
		return errs, existing
	}

	isCorrectPassword := existing.CheckPassword(userDto.Password)

	if !isCorrectPassword || userDto.Email != existing.Email {
		errs = append(errs, rules.ValidationError{
			Field:   "email",
			Message: log.ErrAuthFailed,
		})
	}

	return errs, existing
}

func Login(credentials user_models.Credentials) ([]rules.ValidationError, *user_models.User) {

	validationError, existing := loginValidateCredential(credentials)

	if len(validationError) > 0 {
		return validationError, existing
	}

	return validationError, existing
}
