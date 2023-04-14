package mvm

import (
	"encoding/json"
	"mvm_backend/internal/pkg/generated/mvmPb"
	"mvm_backend/internal/pkg/model"
	"net/http"
)

func (s *MVMServiceServer) GetProfile(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user_id").(string)
	if !ok {
		http.Error(w, "User ID not found", http.StatusInternalServerError)
		return
	}

	res, err := s.service.GetProfile(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mvmPb.GetProfileResponse{Profile: encodeUserProfile(res)})
}

func encodeUserProfile(profile *model.User) *mvmPb.UserProfile {
	return &mvmPb.UserProfile{
		Id:       profile.ID,
		Username: profile.Username,
		Email:    profile.Email,
	}

}
