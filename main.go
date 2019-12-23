package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	messageModel "newland/app/models"
	"newland/bootstrap"
)

func init() {
	bootstrap.MysqlInit()
}

func main() {
	router := gin.Default()
	router.GET("/welcome", func(c *gin.Context) {
		//id := c.Query("id")
		id := 124342
		lastMessage := messageModel.GetLinkMessageOne(c, id)
		c.JSON(http.StatusOK, gin.H{
			"status": 200,
			"error":  nil,
			"data":   lastMessage,
		})
	})
	router.Run(":8998")
}
