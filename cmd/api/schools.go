// Filename: cmd/api/schools.go

package main

import (
	"api/internal/data"
	"fmt"
	"net/http"
	"time"
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
		app.notFoundResponse(w, r)
		return
	}
	// Create a new intance of School struct containing the id we got from the request

	school := data.School{
		ID:        id,
		CreatedAt: time.Now(),
		Name:      "ApplTree",
		Level:     "High School",
		Contact:   "Anna smith",
		Phone:     "601-4411",
		Address:   "14 Apple tree",
		Mode:      []string{"Blended/Hybrid", "Online"},
		Version:   1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"school": school}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}
