package mvm

import (
	"encoding/json"
	"mvm_backend/internal/pkg/generated/mvmPb"
	"mvm_backend/internal/pkg/model"
	"net/http"
)

func (s *MVMServiceServer) SearchForUsers(w http.ResponseWriter, r *http.Request) {
	var input mvmPb.SearchForUsersRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := s.service.SearchForUsers(input.SearchInput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func encodeSearchForUsersResponse(users []*model.User) []*mvmPb.UserProfile {
	out := make([]*mvmPb.UserProfile, 0)
	for i, user := range users {
		out[i] = encodeUserProfile(user)
	}
	return out
}
