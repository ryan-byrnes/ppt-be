package main

import (
	"encoding/json"
	"net/http"

	"github.com/ryan-byrnes/ppt-be/models"
)

func (app *application) getExercise(w http.ResponseWriter, r *http.Request) {

	backsquat := models.Exercise{
		Id:           1,
		ExerciseName: "Back Squat",
	}

	js, err := json.MarshalIndent(backsquat, "", "\t")
	if err != nil {
		app.logger.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}
