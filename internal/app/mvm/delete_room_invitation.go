package mvm

import (
	"encoding/json"
	"mvm_backend/internal/pkg/errors"
	"mvm_backend/internal/pkg/generated/mvmPb"
	"net/http"
)

func (s *MVMServiceServer) DeleteRoomInvitation(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user_id").(string)
	if !ok {
		errors.NewHTTPError(w, errors.NewError("User ID not found", http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	var input mvmPb.DeleteRoomInvitationRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		errors.NewHTTPError(w, errors.NewError(err.Error(), http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if err := s.service.DeleteRoomInvitation(userID, input.RoomId, input.UserId); err != nil {
		errors.NewHTTPError(w, errors.NewError(err.Error(), http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mvmPb.DeleteRoomInvitationResponse{})
}
