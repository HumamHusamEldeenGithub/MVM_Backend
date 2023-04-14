package mvm

import (
	"net/http"
)

func (s *MVMServiceServer) HandleConnections(w http.ResponseWriter, r *http.Request) {
	s.service.HandleConnections(w, r)
}
