package mvm

import (
	"net/http"
)

func (s *MVMServiceServer) HandleWebSocketRTC(w http.ResponseWriter, r *http.Request) {
	s.service.HandleWebSocketRTC(w, r)
}
