package mvm

import (
	"encoding/json"
	"mvm_backend/internal/pkg/errors"
	"mvm_backend/internal/pkg/generated/mvmPb"
	"net/http"
)

func (s *MVMServiceServer) LoginUser(w http.ResponseWriter, r *http.Request) {
	var input mvmPb.LoginUserRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		errors.NewHTTPError(w, errors.NewError(err.Error(), http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	res, err := s.service.LoginUser(input.Username, input.Password)
	if err != nil {

		errors.NewHTTPError(w, errors.NewError(err.Error(), http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mvmPb.LoginUserResponse{
		Token:        res.Token,
		RefreshToken: res.RefreshToken,
	})
}
