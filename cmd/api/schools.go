// Filename: cmd/api/schools.go

package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// createSchoolHandler for the POST reqeust at /v1/schools endpoint
func (app *application) createSchoolHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "create a new school...")
}

// showSchoolHandler for the GET request at /v1/schools/:id endpoint
func (app *application) showSchoolHandler(w http.ResponseWriter, r *http.Request) {
	// Use the params from the context function to get the request context
	params := httprouter.ParamsFromContext(r.Context())

	// Now extract the id from the params
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)

	// If there is an error respond erro not found
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	// Display the school id if no errors are found
	fmt.Fprintf(w, "show the details for the school %d\n", id)
}
