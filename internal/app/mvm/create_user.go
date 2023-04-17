package mvm

import (
	"encoding/json"
	"mvm_backend/internal/pkg/errors"
	"mvm_backend/internal/pkg/generated/mvmPb"
	"mvm_backend/internal/pkg/model"
	"net/http"
)

func (s *MVMServiceServer) CreateUser(w http.ResponseWriter, r *http.Request) {
	var input mvmPb.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		errors.NewHTTPError(w, errors.NewError(err.Error(), http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	userID, err := s.service.CreateUser(&model.User{
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
	})
	if err != nil {
		errors.NewHTTPError(w, errors.NewError(err.Error(), http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mvmPb.CreateUserResponse{
		Id: userID,
	})
}
