package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"backend/schema"
	"backend/service"
)

type FriendshipHandler struct {
	svc *service.FriendshipService
}

func NewFriendshipHandler(svc *service.FriendshipService) *FriendshipHandler {
	return &FriendshipHandler{svc: svc}
}

// POST /api/users/{user_id}/friends
func (h *FriendshipHandler) AddFriend(w http.ResponseWriter, r *http.Request) {
	uIDStr := r.PathValue("user_id")
	userID, err := strconv.Atoi(uIDStr)
	if err != nil || userID <= 0 {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var req schema.AddFriendRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.FriendID <= 0 {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := h.svc.AddFriend(userID, req.FriendID); err != nil {
		if errors.Is(err, service.ErrCannotAddSelf) || errors.Is(err, service.ErrAlreadyFriend) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Friend added successfully",
	})
}

// GET /api/users/{user_id}/friends
func (h *FriendshipHandler) GetFriends(w http.ResponseWriter, r *http.Request) {
	uIDStr := r.PathValue("user_id")
	userID, err := strconv.Atoi(uIDStr)
	if err != nil || userID <= 0 {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	friends, err := h.svc.GetFriends(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(friends)
}
