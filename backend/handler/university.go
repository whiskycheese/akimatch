package handler

import (
	"encoding/json"
	"net/http"
	"backend/service"
)

type UniversityHandler struct {
	svc *service.UniversityService
}

func NewUniversityHandler(svc *service.UniversityService) *UniversityHandler {
	return &UniversityHandler{svc: svc}
}

func (h *UniversityHandler) GetAllUniversities(w http.ResponseWriter, r *http.Request) {
	universities, err := h.svc.GetAllUniversities()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(universities)
}
