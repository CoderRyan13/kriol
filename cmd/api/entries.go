// Filename: cmd/api/entries.go

package main

import (
	"fmt"
	"net/http"
)

// createEntryHandler for the "Post /v1/entries" endpoint
func (app *application) createEntryHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "create a new entry..")
}

// showEntryHandler for the "Get /v1/entries/:id" endpoint
func (app *application) showEntryHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	// Display the entry id
	fmt.Fprintf(w, "show the details for school %d\n", id)
}
