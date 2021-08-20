package controller

import (
	"fmt"
	"github.com/cloverzrg/wechat-work-message-push-go/config"
	"github.com/cloverzrg/wechat-work-message-push-go/qyapi"
	"github.com/gin-gonic/gin"
	"strings"
)

func Index(c *gin.Context) {
	c.String(400, "post to /push")
}

func ServerChan(c *gin.Context) {
	token := c.Param("token.send")
	token = strings.ReplaceAll(token, ".send", "")
	if token != config.Config.Token {
		c.String(400, "invalid token:"+token)
		return
	}
	text := c.Query("text")
	desp := c.PostForm("desp")
	c.String(200, "text:"+text+",desp="+desp)

	msg := fmt.Sprintf("%s\n%s", text, desp)

	err := qyapi.SendMessage(msg, "")
	if err != nil {
		c.String(500, err.Error())
		return
	}
	c.String(200, "ok")
}
