package mvm

import (
	"mvm_backend/internal/pkg/payloads"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *MVMServiceServer) SearchForUsers(c *gin.Context) {
	var input payloads.SearchForUsers
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := s.service.SearchForUsers(input.SearchInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": res})
}
