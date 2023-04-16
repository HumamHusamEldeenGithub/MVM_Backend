package mvm

import (
	"encoding/json"
	"mvm_backend/internal/pkg/generated/mvmPb"
	"net/http"
)

func (s *MVMServiceServer) DeleteRoomInvitation(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user_id").(string)
	if !ok {
		http.Error(w, "User ID not found", http.StatusInternalServerError)
		return
	}
	var input mvmPb.DeleteRoomInvitation
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := s.service.DeleteRoomInvitation(userID, input.RoomId, input.UserId); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mvmPb.Empty{})
}
