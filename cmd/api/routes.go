package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/status", app.statusHandler)

	router.HandlerFunc(http.MethodGet, "/v1/personal-records/:id", app.getPersonalRecords)

	router.HandlerFunc(http.MethodPost, "/v1/add-personal-record", app.addPersonalRecord)

	return router
}
