package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Instanciate a default router which include a logger
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	
	// Define a handler for GET requests
	r.GET("/*path", func(c *gin.Context) {
		path := c.Param("path")

		// conditional message based on the provided path
		if path != "/"{ 
			c.String(200, "hello from %s\n", path)
		} else {
			c.String(200, "try to add /test as path\n")
		}
	})
	// Start the web service on port 8090
	r.Run(":8090")
}