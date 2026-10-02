package repository

import (
	"database/sql"
	"backend/schema"
)

type FacultyRepository struct {
	db *sql.DB
}

func NewFacultyRepository(db *sql.DB) *FacultyRepository {
	return &FacultyRepository{db: db}
}

// 指定した大学IDに紐づく学部一覧を取得する
func (r *FacultyRepository) GetByUniversityID(universityID int) ([]schema.FacultyResponse, error) {
	query := "SELECT id, university_id, name FROM faculties WHERE university_id = $1 ORDER BY id ASC"
	rows, err := r.db.Query(query, universityID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var faculties []schema.FacultyResponse
	for rows.Next() {
		var f schema.FacultyResponse
		if err := rows.Scan(&f.ID, &f.UniversityID, &f.Name); err != nil {
			return nil, err
		}
		faculties = append(faculties, f)
	}

	return faculties, nil
}
