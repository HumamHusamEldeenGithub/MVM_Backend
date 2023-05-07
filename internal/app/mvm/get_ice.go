package mvm

import (
	"encoding/json"
	"mvm_backend/internal/pkg/errors"
	"mvm_backend/internal/pkg/generated/mvmPb"
	"net/http"
)

func (s *MVMServiceServer) GetIce(w http.ResponseWriter, r *http.Request) {
	// userID, ok := r.Context().Value("user_id").(string)
	// if !ok {
	// 	errors.NewHTTPError(w, errors.NewError("User ID not found", http.StatusUnauthorized), http.StatusUnauthorized)
	// 	return
	// }
	userID := r.URL.Query().Get("id")

	Ices, err := s.service.GetIce(userID)
	if err != nil {
		errors.NewHTTPError(w, errors.NewError(err.Error(), http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mvmPb.GetIceResponse{Ices: Ices})
}
