package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"backend/service"
)

type CourseHandler struct {
	svc *service.CourseService
}

func NewCourseHandler(svc *service.CourseService) *CourseHandler {
	return &CourseHandler{svc: svc}
}

func (h *CourseHandler) GetCoursesByUniversityAndFacultyID(w http.ResponseWriter, r *http.Request) {
	// Go 1.22 の PathValue で両方のパスパラメータを取得
	univIDStr := r.PathValue("university_id")
	facIDStr := r.PathValue("faculty_id")

	universityID, err1 := strconv.Atoi(univIDStr)
	facultyID, err2 := strconv.Atoi(facIDStr)

	if err1 != nil || universityID <= 0 || err2 != nil || facultyID <= 0 {
		http.Error(w, "Invalid university ID or faculty ID", http.StatusBadRequest)
		return
	}

	courses, err := h.svc.GetCoursesByUniversityAndFacultyID(universityID, facultyID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}
