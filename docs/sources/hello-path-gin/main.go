package main

import (
	"fmt"
	"log"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Try to add /test as a path.")
	})
	
	r.GET("/:path", func(c *gin.Context) {
		path := c.Param("path")
		c.String(200, "User requested the URL path: %s", path) 
	})

	fmt.Println("Gin GO Hello World example")
	fmt.Println("--> Server running on http://localhost:8090")

	if err := r.Run(":8090"); err != nil {
		log.Fatal(err)
	}
}