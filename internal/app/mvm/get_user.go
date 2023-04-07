package mvm

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *MVMServiceServer) GetUser(c *gin.Context) {
	email := c.Query("email")
	if len(email) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email"})
		return
	}

	user, err := s.service.GetUser(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.Password = ""
	c.JSON(http.StatusOK, *user)
}
