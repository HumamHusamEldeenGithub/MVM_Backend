package mvm

import (
	"encoding/json"
	"mvm_backend/internal/pkg/errors"
	"mvm_backend/internal/pkg/generated/mvmPb"
	"mvm_backend/internal/pkg/model"
	"net/http"
)

func (s *MVMServiceServer) SearchForUsers(w http.ResponseWriter, r *http.Request) {
	var input mvmPb.SearchForUsersRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		errors.NewHTTPError(w, errors.NewError(err.Error(), http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	res, err := s.service.SearchForUsers(input.SearchInput)
	if err != nil {
		errors.NewHTTPError(w, errors.NewError(err.Error(), http.StatusNotFound), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&mvmPb.SearchForUsersResponse{Users: encodeSearchForUsersResponse(res)})
}

func encodeSearchForUsersResponse(users []*model.User) []*mvmPb.UserProfile {
	out := make([]*mvmPb.UserProfile, len(users))
	for i, user := range users {
		out[i] = encodeUserProfile(user)
	}
	return out
}
