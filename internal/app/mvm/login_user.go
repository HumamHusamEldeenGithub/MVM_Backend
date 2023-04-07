package mvm

import (
	"mvm_backend/internal/pkg/payloads"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *MVMServiceServer) LoginUser(c *gin.Context) {
	var input payloads.LoginUserRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := s.service.LoginUser(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, *res)

}
