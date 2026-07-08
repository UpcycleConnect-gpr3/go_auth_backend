package auth_middleware

import (
	"authentication_backend/app/models/user_models"
	"authentication_backend/utils/jwt"
	"authentication_backend/utils/log"
	"context"
	"net/http"
)

type contextKey string

const userIdKey contextKey = "userId"
const roleKey contextKey = "role"

func IsAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, role := jwt.Auth(w, r)

		if userId == "" {
			return
		}

		ctx := context.WithValue(r.Context(), userIdKey, userId)
		ctx = context.WithValue(ctx, roleKey, role)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	}
}

func GetUserId(ctx context.Context) string {

	userId, ok := ctx.Value(userIdKey).(string)

	if !ok {
		log.Info("No user id from context")
		return ""
	}

	return userId
}

func GetRole(ctx context.Context) user_models.Role {

	role, ok := ctx.Value(roleKey).(user_models.Role)

	if !ok {
		log.Info("No role from context")
		return ""
	}

	return role
}
