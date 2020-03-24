package controller

import (
	"github.com/cloverzrg/wechat-work-message-push-go/qyapi"
	"github.com/gin-gonic/gin"
)

func Push(c *gin.Context) {

	msg := c.PostForm("message")
	toUser := c.PostForm("userId")
	err := qyapi.SendMessage(msg, toUser)
	if err != nil {
		c.String(500, err.Error())
		return
	}
	c.String(200, "ok")
}
