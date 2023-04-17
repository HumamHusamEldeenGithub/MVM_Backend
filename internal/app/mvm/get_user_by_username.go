package mvm

import (
	"encoding/json"
	"mvm_backend/internal/pkg/errors"
	"mvm_backend/internal/pkg/generated/mvmPb"
	"net/http"
)

func (s *MVMServiceServer) GetUserByUsername(w http.ResponseWriter, r *http.Request) {
	var input mvmPb.GetUserByUsernameRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		errors.NewHTTPError(w, errors.NewError(err.Error(), http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	res, err := s.service.GetUserByUsername(input.Username)
	if err != nil {
		errors.NewHTTPError(w, errors.NewError(err.Error(), http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&mvmPb.GetUserByUsernameResponse{
		Profile: encodeUserProfile(res),
	})
}
