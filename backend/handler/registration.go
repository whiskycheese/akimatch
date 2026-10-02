package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"backend/service"
	"backend/schema"
)

type RegistrationHandler struct {
	svc *service.RegistrationService
}

func NewRegistrationHandler(svc *service.RegistrationService) *RegistrationHandler {
	return &RegistrationHandler{svc: svc}
}

func (h *RegistrationHandler) GetRegistrationsByUserID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("user_id")
	userID, err := strconv.Atoi(idStr)
	if err != nil || userID <= 0 {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	list, err := h.svc.GetRegistrationsByUserID(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list)
}

func (h *RegistrationHandler) CreateRegistration(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("user_id")
	userID, err := strconv.Atoi(idStr)
	if err != nil || userID <= 0 {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var req schema.CreateRegistrationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// 簡易バリデーション
	if req.CourseID <= 0 || req.Day == "" || req.Period <= 0 {
		http.Error(w, "course_id, day, and period are required", http.StatusBadRequest)
		return
	}

	res, err := h.svc.CreateRegistration(userID, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201 Created
	json.NewEncoder(w).Encode(res)
}

func (h *RegistrationHandler) DeleteRegistration(w http.ResponseWriter, r *http.Request) {
	uIDStr := r.PathValue("user_id")
	rIDStr := r.PathValue("registration_id")

	userID, err1 := strconv.Atoi(uIDStr)
	registrationID, err2 := strconv.Atoi(rIDStr)

	if err1 != nil || userID <= 0 || err2 != nil || registrationID <= 0 {
		http.Error(w, "Invalid user ID or registration ID", http.StatusBadRequest)
		return
	}

	err := h.svc.DeleteRegistration(userID, registrationID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Registration not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Registration deleted successfully",
	})
}
