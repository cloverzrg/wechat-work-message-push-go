package api

import (
	"github.com/cloverzrg/wechat-work-message-push-go/logger"
	"github.com/gin-gonic/gin"
)

func Start() (err error) {
	r := gin.New()
	gin.SetMode(gin.DebugMode)
	r.Use(logger.GinLogger())
	SetRoute(r)
	err = r.Run("0.0.0.0:80")
	return err
}
