package metric_handlers

import (
	"authentication_backend/utils/log"
	"net/http"
)

func Health(w http.ResponseWriter, r *http.Request) {
	log.Api(r)
}
