package mvm

import (
	"encoding/json"
	"mvm_backend/internal/pkg/errors"
	"mvm_backend/internal/pkg/generated/mvmPb"
	"mvm_backend/internal/pkg/model"
	"net/http"
)

func (s *MVMServiceServer) GetProfile(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user_id").(string)
	if !ok {
		errors.NewHTTPError(w, errors.NewError("User ID not found", http.StatusNotFound), http.StatusNotFound)
		return
	}

	params := r.URL.Query()

	// Extract the value of the "user" parameter
	customId := params.Get("user")
	if len(customId) != 0 {
		userID = customId
	}

	profile, err := s.service.GetProfile(userID)
	if err != nil {
		errors.NewHTTPError(w, errors.NewError(err.Error(), http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	rooms, err := s.service.GetUserRooms(userID)
	if err != nil {
		errors.NewHTTPError(w, errors.NewError(err.Error(), http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mvmPb.GetProfileResponse{
		Profile:   encodeUserProfile(profile),
		UserRooms: rooms,
	})
}

func encodeUserProfile(profile *model.User) *mvmPb.UserProfile {
	return &mvmPb.UserProfile{
		Id:       profile.ID,
		Username: profile.Username,
		Email:    profile.Email,
	}

}
