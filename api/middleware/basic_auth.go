package middleware

import (
	"github.com/cloverzrg/wechat-work-message-push-go/config"
	"github.com/gin-gonic/gin"
)

func BasicAuth() func(c *gin.Context) {
	return gin.BasicAuth(gin.Accounts{
		config.Config.GrafanaWebhookUser: config.Config.GrafanaWebhookPassword,
	})
}
