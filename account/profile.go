package account

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetProfile retrieves a user profile.
func GetProfile(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "Not implemented."})
}
