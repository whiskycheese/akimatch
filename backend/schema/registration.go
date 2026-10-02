package schema

type RegistrationResponse struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	CourseID    int    `json:"course_id"`
	Day         string `json:"day"`
	Period      int    `json:"period"`
	SubjectName string `json:"subject_name"` // 追加
	Room        string `json:"room"`         // 追加
}

type CreateRegistrationRequest struct {
	CourseID int    `json:"course_id"`
	Day      string `json:"day"`
	Period   int    `json:"period"`
}
