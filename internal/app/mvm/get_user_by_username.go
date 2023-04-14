package mvm

import (
	"net/http"
)

func (s *MVMServiceServer) GetUserByUsername(w http.ResponseWriter, r *http.Request) {
	// username := c.Query("username")
	// if len(username) == 0 {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username"})
	// 	return
	// }

	// res, err := s.service.GetUserByUsername(username)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }
	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(*res)
}
