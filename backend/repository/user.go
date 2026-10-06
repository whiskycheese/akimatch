package repository

import (
	"database/sql"
	"backend/schema"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetByID(id int) (*schema.UserResponse, error) {
	query := `
		SELECT id, name, share_code, university_id, faculty_id, created_at 
		FROM users 
		WHERE id = $1`

	var u schema.UserResponse
	err := r.db.QueryRow(query, id).Scan(
		&u.ID,
		&u.Name,
		&u.ShareCode,
		&u.UniversityID,
		&u.FacultyID,
		&u.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &u, nil
}
