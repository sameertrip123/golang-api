// Filename: cmd/api/helpers.go

package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) readIDParam(r *http.Request) (int64, error) {
	// Use the params from the context function to get the request context
	params := httprouter.ParamsFromContext(r.Context())

	// Now extract the id from the params
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)

	// If there is an error respond error not found
	if err != nil || id < 1 {
		return 0, errors.New("invalid ID parameter")
	}
	return id, nil
}

func (app *application) writeJSON(w http.ResponseWriter, statusCode int, data interface{}, header http.Header) error {
	// Convert the data into a json object
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}
	// Add a new line
	js = append(js, '\n')

	// Add the headers
	for key, value := range header {
		w.Header()[key] = value
	}

	// Specify in the header that our response will be in json format instead of plain text
	w.Header().Set("Content-Type", "application/json")

	// Write the status code to the header
	w.WriteHeader(statusCode)

	// Write the json byte[] to the json body
	w.Write(js)
	return nil
}
