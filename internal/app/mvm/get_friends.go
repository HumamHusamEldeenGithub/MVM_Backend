package mvm

import (
	"encoding/json"
	"mvm_backend/internal/pkg/errors"
	"mvm_backend/internal/pkg/generated/mvmPb"
	"mvm_backend/internal/pkg/model"
	"net/http"
)

func (s *MVMServiceServer) GetFriends(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user_id").(string)
	if !ok {
		errors.NewHTTPError(w, errors.NewError("User ID not found", http.StatusNotFound), http.StatusNotFound)
		return
	}

	friends, err := s.service.GetFriends(userID)
	if err != nil {
		errors.NewHTTPError(w, errors.NewError(err.Error(), http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var users []*model.User
	if len(friends.Friends) != 0 {
		users, err = s.service.GetProfiles(friends.Friends)
		if err != nil {
			errors.NewHTTPError(w, errors.NewError(err.Error(), http.StatusBadRequest), http.StatusBadRequest)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mvmPb.GetFriendsResponse{
		Profiles:     encodeUserProfiles(users),
		Pending:      friends.Pending,
		SentRequests: friends.Sent,
	})
}

func encodeUserProfiles(profiles []*model.User) []*mvmPb.UserProfile {
	out := make([]*mvmPb.UserProfile, len(profiles))
	for i, profile := range profiles {
		out[i] = encodeUserProfile(profile)
	}
	return out
}
