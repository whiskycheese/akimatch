package repository

import (
	"database/sql"
	"backend/schema"
)

type UniversityRepository struct {
	db *sql.DB
}

func NewUniversityRepository(db *sql.DB) *UniversityRepository {
	return &UniversityRepository{db: db}
}

func (r *UniversityRepository) GetAll() ([]schema.UniversityResponse, error) {
	rows, err := r.db.Query("SELECT id, name FROM universities ORDER BY id ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var universities []schema.UniversityResponse
	for rows.Next() {
		var u schema.UniversityResponse
		if err := rows.Scan(&u.ID, &u.Name); err != nil {
			return nil, err
		}
		universities = append(universities, u)
	}

	return universities, nil
}
