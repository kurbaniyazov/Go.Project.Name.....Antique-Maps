package main

import (
	"errors"
	"fmt"
	"maps.alexedwards.net/internal/data"
	"maps.alexedwards.net/internal/validator"
	"net/http"
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

	antiqueMaps := &data.AntiqueMaps{
		Title:     input.Title,
		Year:      input.Year,
		Country:   input.Country,
		Condition: input.Condition,
		Type:      input.Type,
	}

	v := validator.New()

	if data.ValidateAntiqueMaps(v, antiqueMaps); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	err = app.models.AntiqueMaps.Insert(antiqueMaps)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/antique-maps/%d", antiqueMaps.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"antique map": antiqueMaps}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) showAntiqueMapHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	//fmt.Fprintf(w, "show the details of antique map %d\n", id)

	antiqueMap, err := app.models.AntiqueMaps.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"antiqueMaps": antiqueMap}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}
