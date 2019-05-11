package account

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RequireAuth returns an unauthorized error if the user is not logged in.
func RequireAuth(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "Not implemented."})
}

// Login authenticates the user and starts a session.
func Login(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "Not implemented."})
}

// Logout ends the users session.
func Logout(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "Not implemented."})
}

// test
