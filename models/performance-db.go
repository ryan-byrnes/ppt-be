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

func (m *DBModel) Get(id int) ([]*PersonalRecord, error) {

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
