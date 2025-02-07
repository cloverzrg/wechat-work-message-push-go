package middleware

import (
	"fmt"
	"github.com/cloverzrg/wechat-work-message-push-go/config"
	"github.com/gin-gonic/gin"
)

func BasicAuth() func(c *gin.Context) {
	if config.Config.GrafanaWebhookUser == "" {
		return func(c *gin.Context) {
			_ = c.AbortWithError(500, fmt.Errorf("Grafana Webhook User not set in configuration. Please check your config file. "))
			return
		}
	}
	return gin.BasicAuth(gin.Accounts{
		config.Config.GrafanaWebhookUser: config.Config.GrafanaWebhookPassword,
	})
}
