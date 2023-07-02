package mvm

import (
	"encoding/json"
	"mvm_backend/internal/pkg/errors"
	"mvm_backend/internal/pkg/generated/mvmPb"
	"mvm_backend/internal/pkg/model"
	"net/http"
)

func (s *MVMServiceServer) GetPendingFriends(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user_id").(string)
	if !ok {
		errors.NewHTTPError(w, errors.NewError("User ID not found", http.StatusNotFound), http.StatusNotFound)
		return
	}

	usersIds, err := s.service.GetPendingFriends(userID)
	if err != nil {
		errors.NewHTTPError(w, errors.NewError(err.Error(), http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var users []*model.User
	if len(usersIds) != 0 {
		users, err = s.service.GetProfiles(usersIds)
		if err != nil {
			errors.NewHTTPError(w, errors.NewError(err.Error(), http.StatusBadRequest), http.StatusBadRequest)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mvmPb.GetPendingFriendsResponse{Profiles: encodeUserProfiles(users)})
}
