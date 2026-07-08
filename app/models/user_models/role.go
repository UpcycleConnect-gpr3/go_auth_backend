package user_models

type Role string

const (
	RoleProfessional  Role = "professional"
	RoleProvider      Role = "provider"
	RoleCreator       Role = "creator"
	RoleEmployee      Role = "employee"
	RoleAdministrator Role = "administrator"
)

func AllUserSelectableRoles() []Role {

	return []Role{RoleProvider, RoleProfessional}
}

func AllAdminAssignableRoles() []Role {
	return []Role{RoleCreator, RoleEmployee, RoleAdministrator}
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
