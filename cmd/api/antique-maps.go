package main

import (
	"fmt"
	"net/http"
)

func (app *application) createAntiqueMapHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new antique map")
}

func (app *application) showAntiqueMapHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "show the details of antique map %d\n", id)
}
