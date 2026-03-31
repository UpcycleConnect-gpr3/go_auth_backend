package source_middleware

import (
	"authentication_backend/utils/log"
	"net/http"
)

func Container(allowedContainer string) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			clientContainer := r.Header.Get("X-Container-Name")
			if clientContainer == "" {
				log.ApiCodeStatus(w, http.StatusForbidden, "", nil)
				return
			}

			if clientContainer != allowedContainer {
				log.ApiCodeStatus(w, http.StatusForbidden, "", nil)
				return
			}
			next.ServeHTTP(w, r)
		}
	}
}
