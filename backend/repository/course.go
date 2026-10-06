package repository

import (
	"database/sql"
	"backend/schema"
)

type CourseRepository struct {
	db *sql.DB
}

func NewCourseRepository(db *sql.DB) *CourseRepository {
	return &CourseRepository{db: db}
}

// university_id と faculty_id の両方を検証して講義一覧を取得
func (r *CourseRepository) GetByUniversityAndFacultyID(universityID, facultyID int) ([]schema.CourseResponse, error) {
	query := `
		SELECT id, university_id, faculty_id, subject_name, COALESCE(room, '') 
		FROM courses 
		WHERE university_id = $1 AND faculty_id = $2 
		ORDER BY id ASC`
	
	rows, err := r.db.Query(query, universityID, facultyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []schema.CourseResponse
	for rows.Next() {
		var c schema.CourseResponse
		if err := rows.Scan(&c.ID, &c.UniversityID, &c.FacultyID, &c.SubjectName, &c.Room); err != nil {
			return nil, err
		}
		courses = append(courses, c)
	}

	return courses, nil
}
