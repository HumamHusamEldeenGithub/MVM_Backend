package mvm

import (
	"encoding/json"
	"mvm_backend/internal/pkg/errors"
	"mvm_backend/internal/pkg/generated/mvmPb"
	"net/http"
)

func (s *MVMServiceServer) GetChat(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user_id").(string)
	if !ok {
		errors.NewHTTPError(w, errors.NewError("User ID not found", http.StatusNotFound), http.StatusNotFound)
		return
	}
	var input mvmPb.GetChatRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		errors.NewHTTPError(w, errors.NewError(err.Error(), http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	chat, err := s.service.GetChat(userID, input.UserId)
	if err != nil {
		errors.NewHTTPError(w, errors.NewError(err.Error(), http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mvmPb.GetChatResponse{Chat: chat})
}
