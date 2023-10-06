package main

import (
	"fmt"
	"maps.alexedwards.net/internal/data"
	"maps.alexedwards.net/internal/validator"
	"net/http"
	"time"
)

func (app *application) createAntiqueMapHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title     string `json:"title"`
		Year      int32  `json:"year"`
		Country   string `json:"country"`
		Condition string `json:"condition"`
		Type      string `json:"type"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	maps := &data.Maps{
		Title:     input.Title,
		Year:      input.Year,
		Country:   input.Country,
		Condition: input.Condition,
		Type:      input.Type,
	}
	// Initialize a new Validator.
	v := validator.New()
	// Call the ValidateMovie() function and return a response containing the errors if
	// any of the checks fail.
	if data.ValidateMaps(v, maps); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
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
