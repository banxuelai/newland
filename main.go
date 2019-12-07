package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/welcome", func(c *gin.Context) {
		name := c.Query("name")
		// fmt.Println("Hello %s", name)
		c.JSON(http.StatusOK, gin.H{
			"status": 200,
			"error":  nil,
			"data":   name,
		})
	})
	router.Run(":8998")
}
