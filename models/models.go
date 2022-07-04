package models

import "time"

type Exercise struct {
	Id           int    `json:"id"`
	ExerciseName string `json:"exercise_name"`
}

type PersonalRecord struct {
	Id         int       `json:"id"`
	ExerciseId int       `json:"exercise_id"`
	Exercise   Exercise  `json:"exercise"`
	Reps       int       `json:"reps"`
	Weight     int       `json:"weight"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
