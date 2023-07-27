package mvm

import (
	"encoding/json"
	"mvm_backend/internal/pkg/errors"
	"mvm_backend/internal/pkg/generated/mvmPb"
	"net/http"
)

func (s *MVMServiceServer) GetUserProfileFeatures(w http.ResponseWriter, r *http.Request) {
	_, ok := r.Context().Value("user_id").(string)
	if !ok {
		errors.NewHTTPError(w, errors.NewError("User ID not found", http.StatusNotFound), http.StatusNotFound)
		return
	}

	var input mvmPb.GetUserProfileFeaturesRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		errors.NewHTTPError(w, errors.NewError(err.Error(), http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	profile, err := s.service.GetProfile(input.Id)
	if err != nil {
		errors.NewHTTPError(w, errors.NewError(err.Error(), http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mvmPb.GetUserProfileFeaturesResponse{
		Profile:        encodeUserProfile(profile),
		AvatarSettings: profile.AvatarSettings,
	})
}
