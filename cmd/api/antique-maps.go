package main

import (
	"fmt"
	"maps.alexedwards.net/internal/data"
	"net/http"
	"time"
)

func (app *application) createAntiqueMapHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title     string    `json:"title"`
		Year      data.Year `json:"year"`
		Country   string    `json:"country"`
		Condition string    `json:"condition"`
		Type      string    `json:"type"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) showAntiqueMapHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
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
		app.serverErrorResponse(w, r, err)
	}

}
