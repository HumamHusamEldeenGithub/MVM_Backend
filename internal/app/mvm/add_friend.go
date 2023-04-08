package mvm

import (
	"mvm_backend/internal/pkg/jwt_manager"
	"mvm_backend/internal/pkg/payloads"
	"mvm_backend/internal/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *MVMServiceServer) AddFriend(c *gin.Context) {
	userID, err := jwt_manager.GetUserIDFromToken(utils.GetJWTToken(c))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var input payloads.AddFriendRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := s.service.AddFriend(userID, input.FriendID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}
