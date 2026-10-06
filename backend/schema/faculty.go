package schema

type FacultyResponse struct {
	ID           int    `json:"id"`
	UniversityID int    `json:"university_id"`
	Name         string `json:"name"`
}
