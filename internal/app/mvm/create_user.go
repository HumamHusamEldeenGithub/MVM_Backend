package mvm

import (
	"mvm_backend/internal/pkg/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *MVMServiceServer) CreateUser(c *gin.Context) {
	var input model.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := s.service.CreateUser(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": userID})
}
