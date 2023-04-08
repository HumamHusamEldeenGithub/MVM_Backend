package mvm

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *MVMServiceServer) GetUserByUsername(c *gin.Context) {
	username := c.Query("username")
	if len(username) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username"})
		return
	}

	user, err := s.service.GetUserByUsername(username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, *user)
}
