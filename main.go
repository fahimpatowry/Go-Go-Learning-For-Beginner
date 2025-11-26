package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	r := gin.Default() // create a router

	// simple GET API
	r.GET("/user", func(c *gin.Context) {
		user := map[string]interface{}{
			"id":   1,
			"name": "Fahim",
			"age":  25,
		}
		c.JSON(http.StatusOK, user)
	})

	// run server on port 8080
	r.Run(":8080")
}