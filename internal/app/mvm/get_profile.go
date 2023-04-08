package mvm

import (
	"mvm_backend/internal/pkg/jwt_manager"
	"mvm_backend/internal/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *MVMServiceServer) GetProfile(c *gin.Context) {
	userID, err := jwt_manager.GetUserIDFromToken(utils.GetJWTToken(c))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := s.service.GetProfile(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, *user)
}
