// The rap-server command stands up a RAP server.
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mccormickbrand/rap-server/account"
)

// main sets up the server router and starts the server.
func main() {
	router := gin.Default()
	auth := router.Group("/auth")
	auth.POST("/login", account.Login)
	auth.Any("/logout", account.Logout)
	secure := router.Group("/secure")
	// bind profile endpoints
	secure.GET("/profile", account.GetProfile)
	router.Run()
}
