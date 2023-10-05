package main

import (
	"fmt"
	"maps.alexedwards.net/internal/data"
	"net/http"
	"time"
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
	//fmt.Fprintf(w, "show the details of antique map %d\n", id)

	maps := data.Maps{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Italy Map",
		Year:      1843,
		Country:   "Italy",
		Condition: "Well",
		Type:      "Exploration Map",
		Version:   1,
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"maps": maps}, nil)
	if err != nil {
		app.Logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}

}
