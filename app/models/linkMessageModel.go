package messageModel

import (
	"github.com/gin-gonic/gin"
	"newland/app/models/message"
	"newland/bootstrap"
)

func GetLinkMessageOne(c *gin.Context,id int) []message.LinkMessage {
	var data []message.LinkMessage
	err := bootstrap.MysqlEngine["message"].Table("link_message").
		Where("id = ?", id).
		Find(&data)
	if err != nil {
		panic(err)
	}
	return  data
}
