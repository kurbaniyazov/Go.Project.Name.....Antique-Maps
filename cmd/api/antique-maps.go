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

func (app *application) updateAntiqueMapsHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	antiqueMaps, err := app.models.AntiqueMaps.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return

	}
	var input struct {
		Title     *string `json:"title"`
		Year      *int32  `json:"year"`
		Country   *string `json:"country"`
		Condition *string `json:"condition"`
		Type      *string `json:"type"`
	}
	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if input.Title != nil {
		antiqueMaps.Title = *input.Title
	}
	if input.Year != nil {
		antiqueMaps.Year = *input.Year
	}
	if input.Country != nil {
		antiqueMaps.Country = *input.Country
	}
	if input.Condition != nil {
		antiqueMaps.Condition = *input.Condition
	}
	if input.Type != nil {
		antiqueMaps.Type = *input.Type
	}

	v := validator.New()
	if data.ValidateAntiqueMaps(v, antiqueMaps); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.AntiqueMaps.Update(antiqueMaps)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrEditConflict):
			app.editConflictResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"antique maps": antiqueMaps}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deleteAntiqueMapsHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the movie ID from the URL.
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	// Delete the movie from the database, sending a 404 Not Found response to the
	// client if there isn't a matching record.
	err = app.models.AntiqueMaps.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	// Return a 200 OK status code along with a success message.
	err = app.writeJSON(w, http.StatusOK, envelope{"message": "antique map successfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) listAntiqueMapsHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title   string
		Country string
		data.Filters
	}
	v := validator.New()
	qs := r.URL.Query()
	input.Title = app.readString(qs, "title", "")
	input.Country = app.readString(qs, "country", "")

	input.Filters.Page = app.readInt(qs, "page", 1, v)
	input.Filters.PageSize = app.readInt(qs, "page_size", 20, v)

	input.Filters.Sort = app.readString(qs, "sort", "id")
	input.Filters.SortSafelist = []string{"id", "title", "year", "country", "-id", "-title", "-year", "-country"}

	if data.ValidateFilters(v, input.Filters); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	antiqueMapss, metadata, err := app.models.AntiqueMaps.GetAll(input.Title, input.Country, input.Filters)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"antique maps": antiqueMapss, "metadata": metadata}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
