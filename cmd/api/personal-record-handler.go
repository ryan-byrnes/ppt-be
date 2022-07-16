package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/ryan-byrnes/ppt-be/models"
)

func (app *application) getPersonalRecords(w http.ResponseWriter, r *http.Request) {

	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.logger.Print(errors.New("invalid userId"))
		app.errorJson(w, err)
		return
	}

	personalRecords, _ := app.models.DB.GetAllPersonalRecords(id)

	app.logger.Println("personal records", personalRecords)

	app.writeJson(w, http.StatusOK, personalRecords, "personal records")
}

type personalRecordPayload struct {
	UserId       string `json:"user_id"`
	ExerciseId   string `json:"exercise_id"`
	ExerciseName string `json:"exercise_name"`
	Reps         string `json:"reps"`
	Weight       string `json:"weight"`
}

func (app *application) addPersonalRecord(w http.ResponseWriter, r *http.Request) {
	var payload personalRecordPayload

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		app.logger.Println(err)
		app.errorJson(w, err)
		return
	}

	var personalRecord models.PersonalRecord

	personalRecord.UserId, _ = strconv.Atoi(payload.UserId)
	personalRecord.ExerciseId, _ = strconv.Atoi(payload.ExerciseId)
	personalRecord.ExerciseName = payload.ExerciseName
	personalRecord.Reps, _ = strconv.Atoi(payload.Reps)
	personalRecord.Weight, _ = strconv.Atoi(payload.Weight)
	personalRecord.CreatedAt = time.Now()
	personalRecord.UpdatedAt = time.Now()

	type jsonResp struct {
		OK bool `json:"ok"`
	}

	ok := jsonResp{
		OK: true,
	}

	err = app.writeJson(w, http.StatusOK, ok, "response")
	if err != nil {
		app.errorJson(w, err)
		return
	}

	err = app.models.DB.InsertPersonalRecord(personalRecord)
	if err != nil {
		app.errorJson(w, err)
		return
	}

}
