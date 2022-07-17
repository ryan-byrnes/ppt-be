package models

import (
	"context"
	"database/sql"
	"log"
	"time"
)

type DBModel struct {
	DB *sql.DB
}

func (m *DBModel) GetAllPersonalRecords(id int) ([]*PersonalRecord, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select "exerciseName", "reps", "weight" from "personalRecords" where "userId" = $1`

	rows, _ := m.DB.QueryContext(ctx, query, id)
	defer rows.Close()

	var personalRecords []*PersonalRecord

	for rows.Next() {
		var personalRecord PersonalRecord

		err := rows.Scan(
			&personalRecord.ExerciseName,
			&personalRecord.Reps,
			&personalRecord.Weight,
		)
		if err != nil {
			return nil, err
		}
		personalRecords = append(personalRecords, &personalRecord)
	}

	log.Println("personal Records: ", personalRecords)

	return personalRecords, nil
}

func (m *DBModel) InsertPersonalRecord(personalRecord PersonalRecord) error {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `insert into "personalRecords" (user_id, exercise_id, exercise_name, reps, weight, created_at, updated_at) 
			  values ($1, $2, $3, $4, $5, $6, $7))`

	_, err := m.DB.ExecContext(ctx, query,
		personalRecord.UserId,
		personalRecord.ExerciseId,
		personalRecord.ExerciseName,
		personalRecord.Reps,
		personalRecord.Weight,
		personalRecord.CreatedAt,
		personalRecord.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil

}

func (m *DBModel) UpdatePersonalRecord(personalRecord PersonalRecord) error {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `update "personalRecords" 
	set reps = $1,
		weight = $2,
		updated_at = $3
	where exercise_id = $4`

	_, err := m.DB.ExecContext(ctx, query,
		personalRecord.Reps,
		personalRecord.Weight,
		personalRecord.UpdatedAt,
		personalRecord.ExerciseId)
	if err != nil {
		return err
	}

	return nil
}
