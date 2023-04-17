package mvm

import (
	"encoding/json"
	"mvm_backend/internal/pkg/errors"
	"mvm_backend/internal/pkg/generated/mvmPb"
	"net/http"
)

func (s *MVMServiceServer) GetFriends(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user_id").(string)
	if !ok {
		errors.NewHTTPError(w, errors.NewError("User ID not found", http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	usersIds, err := s.service.GetFriends(userID)
	if err != nil {
		errors.NewHTTPError(w, errors.NewError(err.Error(), http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mvmPb.GetFriendsReposnes{Ids: usersIds})
}
