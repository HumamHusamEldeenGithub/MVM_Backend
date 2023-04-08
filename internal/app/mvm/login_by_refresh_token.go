package mvm

import (
	"mvm_backend/internal/pkg/payloads"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *MVMServiceServer) LoginByRefreshToken(c *gin.Context) {
	var input payloads.LoginByRefreshToken
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := s.service.LoginByRefreshToken(input.RefreshToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, *res)
}
