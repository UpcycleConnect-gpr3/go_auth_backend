package user_models

type Role string

const (
	RoleProfessional  Role = "professional"
	RoleProvider      Role = "provider"
	RoleCreator       Role = "creator"
	RoleEmployee      Role = "employee"
	RoleAdministrator Role = "administrator"
)

// Rôles sélectionnables par l'utilisateur à l'inscription (frontend).
func AllUserSelectableRoles() []Role {
	return []Role{RoleProfessional, RoleProvider, RoleCreator}
}

// Rôles assignables uniquement par un administrateur (backoffice).
func AllAdminAssignableRoles() []Role {
	return []Role{RoleEmployee, RoleAdministrator}
}

func AllRoles() []Role {
	return append(AllUserSelectableRoles(), AllAdminAssignableRoles()...)
}

func IsValidRole(role string) bool {
	for _, r := range AllRoles() {
		if string(r) == role {
			return true
		}
	}
	return false
}

func (r Role) IsUserSelectable() bool {
	for _, role := range AllUserSelectableRoles() {
		if role == r {
			return true
		}
	}
	return false
}

func (r Role) IsAdminAssignable() bool {
	for _, role := range AllAdminAssignableRoles() {
		if role == r {
			return true
		}
	}
	return false
}
