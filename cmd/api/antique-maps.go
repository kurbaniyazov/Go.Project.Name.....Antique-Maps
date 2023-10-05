package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func (app *application) createAntiqueMapHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new antique map")
}

func (app *application) showAntiqueMapHandler(w http.ResponseWriter, r *http.Request) {

	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "show the details of antique map %d\n", id)
}
