package main

import (
	"net/http"

	"github.com/ryan-byrnes/ppt-be/models"
)

func (app *application) getExercise(w http.ResponseWriter, r *http.Request) {

	backsquat := models.Exercise{
		Id:           1,
		ExerciseName: "Back Squat",
	}

	app.writeJson(w, http.StatusOK, backsquat, "exercise")
}
