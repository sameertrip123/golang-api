// Filename: cmd/api/routes.go

package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	// Create a new http router instance
	router := httprouter.New()
	// Handling the various routes
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/schools", app.createSchoolHandler)
	router.HandlerFunc(http.MethodGet, "/v1/school:id", app.showSchoolHandler)
	return router
}
