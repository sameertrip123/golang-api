// Filename: cmd/api/errors.go
package main

import (
	"fmt"
	"net/http"
)

func (app *application) logError(r *http.Request, err error) {
	app.logger.Println(err)
}

func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, statusCode int, message interface{}) {
	// Create the JSON response
	env := envelope{"error": message}

	err := app.writeJSON(w, statusCode, env, nil)

	if err != nil {
		app.logError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// Server error messages
func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)

	// Prepare the error with a message
	message := "the server encountered an error and could not process the request"
	app.errorResponse(w, r, http.StatusInternalServerError, message)
}

// Not found error response
func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	app.errorResponse(w, r, http.StatusNotFound, message)
}

// A method not allowed error response
func (app *application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)

	app.errorResponse(w, r, http.StatusMethodNotAllowed, message)
}
