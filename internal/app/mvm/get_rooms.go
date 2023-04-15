package mvm

import (
	"encoding/json"
	"mvm_backend/internal/pkg/generated/mvmPb"
	"net/http"
)

func (s *MVMServiceServer) GetRooms(w http.ResponseWriter, r *http.Request) {
	_, ok := r.Context().Value("user_id").(string)
	if !ok {
		http.Error(w, "User ID not found", http.StatusInternalServerError)
		return
	}

	rooms, err := s.service.GetRooms()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mvmPb.GetRoomsResponse{Rooms: rooms})
}
