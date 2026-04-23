package {{.PackageName}}

import (
	"authentication_backend/utils/log"
	"authentication_backend/utils/response"
	"net/http"
)

func Index{{.ResourceName}}Handler(w http.ResponseWriter, r *http.Request) {
	log.Api(r)
	// TODO: Implement
	response.NewSuccessMessage(w, "OK")
}

func Store{{.ResourceName}}Handler(w http.ResponseWriter, r *http.Request) {
	log.Api(r)
	// TODO: Implement
	response.NewSuccessMessage(w, "OK")
}

func Show{{.ResourceName}}Handler(w http.ResponseWriter, r *http.Request) {
	log.Api(r)
	// TODO: Implement
	response.NewSuccessMessage(w, "OK")
}

func Update{{.ResourceName}}Handler(w http.ResponseWriter, r *http.Request) {
	log.Api(r)
	// TODO: Implement
	response.NewSuccessMessage(w, "OK")
}

func Delete{{.ResourceName}}Handler(w http.ResponseWriter, r *http.Request) {
	log.Api(r)
	// TODO: Implement
	response.NewSuccessMessage(w, "OK")
}
