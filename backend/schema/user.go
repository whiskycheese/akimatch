package schema

import "time"

type UserResponse struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	ShareCode    string    `json:"share_code"`
	UniversityID int       `json:"university_id"`
	FacultyID    int       `json:"faculty_id"`
	CreatedAt    time.Time `json:"created_at"`
}
