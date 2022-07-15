package models

import (
	"context"
	"database/sql"
	"time"
)

type DBModel struct {
	DB *sql.DB
}

func (m *DBModel) Get(id int) ([]*PersonalRecord, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select exercise, reps, weight, createdAt from personalRecords where id = $1`

	rows, _ := m.DB.QueryContext(ctx, query, id)
	defer rows.Close()

	var personalRecords []*PersonalRecord

	for rows.Next() {
		var personalRecord PersonalRecord

		err := rows.Scan(
			&personalRecord.Exercise,
			&personalRecord.Reps,
			&personalRecord.Weight,
			&personalRecord.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		personalRecords = append(personalRecords, &personalRecord)
	}

	return personalRecords, nil
}
