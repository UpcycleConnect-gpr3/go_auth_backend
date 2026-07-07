package user_actions

import (
	"authentication_backend/app/models/user_models"
	"authentication_backend/utils/rules"
	"authentication_backend/var/database"
)

func createValidateUser(userDto user_models.RegisterRequest) []rules.ValidationError {
	var errs []rules.ValidationError

	rules.StringMinLength(userDto.Email, 5, "email", &errs)
	rules.StringMinLength(userDto.Password, 6, "password", &errs)
	rules.StringMaxLength(userDto.Password, 128, "password", &errs)
	rules.MustContainsAny(userDto.Password, "!@#$%^&*()", 1, "password", &errs)

	if !user_models.IsValidRole(string(userDto.Role)) {
		errs = append(errs, rules.ValidationError{
			Field:   "role",
			Message: "Invalid role",
		})
	} else if !userDto.Role.IsUserSelectable() {
		errs = append(errs, rules.ValidationError{
			Field:   "role",
			Message: "This role cannot be selected during registration",
		})
	}

	rules.Unique[user_models.User](database.Auth, user_models.TABLE, "email", "email = ?", &errs, userDto.Email)

	return errs
}

func CreateUser(userDto user_models.RegisterRequest) (*user_models.User, []rules.ValidationError) {

	// Rétrocompatibilité : les clients qui n'envoient pas encore de rôle
	// obtiennent le rôle par défaut plutôt qu'une erreur de validation.
	if userDto.Role == "" {
		userDto.Role = user_models.RoleProfessional
	}

	validationError := createValidateUser(userDto)

	if len(validationError) > 0 {
		return nil, validationError
	}

	user := &user_models.User{}

	if err := user.Create(userDto.Credentials, userDto.Role); err != nil {
		return nil, []rules.ValidationError{{
			Field:   "email",
			Message: "An error has occurred",
		}}
	}

	return user, nil
}
