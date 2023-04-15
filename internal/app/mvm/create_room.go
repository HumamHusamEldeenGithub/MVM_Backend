package mvm

import (
	"encoding/json"
	"mvm_backend/internal/pkg/generated/mvmPb"
	"net/http"
)

func (s *MVMServiceServer) CreateRoom(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user_id").(string)
	if !ok {
		http.Error(w, "User ID not found", http.StatusUnauthorized)
		return
	}

	var input mvmPb.Room
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	input.OwnerId = userID

	room, err := s.service.CreateRoom(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mvmPb.CreateRoomResponse{Room: room})
}
