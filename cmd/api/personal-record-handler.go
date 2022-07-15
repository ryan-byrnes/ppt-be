package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) getPersonalRecords(w http.ResponseWriter, r *http.Request) {

	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.logger.Print(errors.New("invalid id"))
		app.errorJson(w, err)
		return
	}

	personalRecords, err := app.models.DB.Get(id)

	app.writeJson(w, http.StatusOK, personalRecords, "personal records")
}
