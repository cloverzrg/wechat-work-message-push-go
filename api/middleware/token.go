package middleware

import (
	"github.com/cloverzrg/wechat-work-message-push-go/config"
	"github.com/gin-gonic/gin"
)

func TokenMiddleware(c *gin.Context) {
	token := c.GetHeader("token")
	if token == "" {
		token = c.Query("token")
	}
	if token != config.Config.Token {
		c.String(401, "token 错误")
		c.Abort()
		return
	}
}
