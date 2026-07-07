package auth_middleware

import (
	"authentication_backend/app/models/user_models"
	"authentication_backend/utils/log"
	"authentication_backend/utils/response"
	"net/http"
)

// RequireRoles n'autorise l'accès qu'aux rôles listés (lus depuis le JWT).
func RequireRoles(requiredRoles ...user_models.Role) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			role := GetRole(r.Context())

			if role == "" {
				response.NewErrorMessage(w, response.ErrAuthTokenRequired, http.StatusUnauthorized)
				return
			}

			for _, requiredRole := range requiredRoles {
				if role == requiredRole {
					next.ServeHTTP(w, r)
					return
				}
			}

			log.Info("Role '" + string(role) + "' does not have access to this route")
			response.NewErrorMessage(w, response.ErrForbidden, http.StatusForbidden)
		}
	}
}

func RequireAdmin(next http.HandlerFunc) http.HandlerFunc {
	return RequireRoles(user_models.RoleAdministrator)(next)
}
