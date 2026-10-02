package schema

type CourseResponse struct {
	ID           int    `json:"id"`
	UniversityID int    `json:"university_id"`
	FacultyID    int    `json:"faculty_id"`
	SubjectName  string `json:"subject_name"`
	Room         string `json:"room"`
}
