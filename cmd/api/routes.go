package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)

	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	router.HandlerFunc(http.MethodGet, "/v1/antique-maps", app.listAntiqueMapsHandler)
	router.HandlerFunc(http.MethodPost, "/v1/antique-maps", app.createAntiqueMapHandler)
	router.HandlerFunc(http.MethodGet, "/v1/antique-maps/:id", app.showAntiqueMapHandler)

	router.HandlerFunc(http.MethodPatch, "/v1/antique-maps/:id", app.updateAntiqueMapsHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/antique-maps/:id", app.deleteAntiqueMapsHandler)

	router.HandlerFunc(http.MethodPost, "/v1/users", app.registerUserHandler)
	router.HandlerFunc(http.MethodPut, "/v1/users/activated", app.activateUserHandler)

	router.HandlerFunc(http.MethodPost, "/v1/tokens/authentication", app.createAuthenticationTokenHandler)

	return app.recoverPanic(app.rateLimit(router))
}
