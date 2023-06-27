// Filename: cmd/api/schools.go

package main

import (
	"fmt"
	"net/http"
)

// createSchoolHandler for the POST request at /v1/schools endpoint
func (app *application) createSchoolHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "create a new school...")
}

// showSchoolHandler for the GET request at /v1/schools/:id endpoint
func (app *application) showSchoolHandler(w http.ResponseWriter, r *http.Request) {
	// Get the id and the err from the readIDParam function by passing the request object
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	// Display the school id if no errors are found
	fmt.Fprintf(w, "show the details for the school %d\n", id)
}
