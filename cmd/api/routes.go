package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/antique-maps", app.createAntiqueMapHandler)
	router.HandlerFunc(http.MethodGet, "/v1/antique-maps/:id", app.showAntiqueMapHandler)

	return router
}
