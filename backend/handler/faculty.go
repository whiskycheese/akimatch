package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"backend/service"
)

type FacultyHandler struct {
	svc *service.FacultyService
}

func NewFacultyHandler(svc *service.FacultyService) *FacultyHandler {
	return &FacultyHandler{svc: svc}
}

func (h *FacultyHandler) GetFacultiesByUniversityID(w http.ResponseWriter, r *http.Request) {
	// URL パスの {id} 部分を取得（例: /api/universities/1/faculties -> "1"）
	idStr := r.PathValue("id")
	universityID, err := strconv.Atoi(idStr)
	if err != nil || universityID <= 0 {
		http.Error(w, "Invalid university ID", http.StatusBadRequest)
		return
	}

	faculties, err := h.svc.GetFacultiesByUniversityID(universityID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(faculties)
}
