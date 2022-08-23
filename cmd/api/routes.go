package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/status", app.statusHandler)

	router.HandlerFunc(http.MethodGet, "/v1/personal-records/:id", app.getPersonalRecords)

	router.HandlerFunc(http.MethodPost, "/v1/add-personal-record", app.addPersonalRecord)

	router.HandlerFunc(http.MethodPut, "/v1/update-personal-record/:exerciseId", app.updatePersonalRecord)

	return app.enableCORS(router)
}
