package mvm

import (
	"encoding/json"
	"mvm_backend/internal/pkg/generated/mvmPb"
	"net/http"
)

func (s *MVMServiceServer) LoginByRefreshToken(w http.ResponseWriter, r *http.Request) {
	var input mvmPb.LoginByRefreshToken
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := s.service.LoginByRefreshToken(input.RefreshToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mvmPb.LoginUserResponse{
		Token:        res.Token,
		RefreshToken: res.RefreshToken,
	})
}
