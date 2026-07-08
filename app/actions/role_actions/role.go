package role_actions

import (
	"authentication_backend/app/models/user_models"
	"authentication_backend/utils/rules"
)

func UpdateUserRole(adminId string, targetUserId string, newRole user_models.Role) []rules.ValidationError {
	var errs []rules.ValidationError

	if !user_models.IsValidRole(string(newRole)) {
		errs = append(errs, rules.ValidationError{
			Field:   "role",
			Message: "Invalid role",
		})
		return errs
	}

	var admin user_models.User
	if err := admin.Get([]string{"id", "role"}, "id = ?", adminId); err != nil {
		errs = append(errs, rules.ValidationError{
			Field:   "admin",
			Message: "Admin not found",
		})
		return errs
	}

	if admin.Role != user_models.RoleAdministrator {
		errs = append(errs, rules.ValidationError{
			Field:   "admin",
			Message: "Only administrators can assign roles",
		})
		return errs
	}

	var targetUser user_models.User
	if err := targetUser.Get([]string{"id", "role"}, "id = ?", targetUserId); err != nil {
		errs = append(errs, rules.ValidationError{
			Field:   "user",
			Message: "User not found",
		})
		return errs
	}

	if err := user_models.UpdateUserRole(targetUserId, newRole); err != nil {
		errs = append(errs, rules.ValidationError{
			Field:   "role",
			Message: "Failed to update role",
		})
	}

	return errs
}
