package api

import (
	"github.com/cloverzrg/wechat-work-message-push-go/api/controller"
	"github.com/cloverzrg/wechat-work-message-push-go/api/middleware"
	"github.com/gin-gonic/gin"
)

func SetRoute(r *gin.Engine) {
	r.GET("/", controller.Index)
	r.POST("/push", middleware.TokenMiddleware, controller.Push)
	r.GET("/push", middleware.TokenMiddleware, controller.PushInGet)
	r.POST("/grafana", middleware.BasicAuth(), controller.GrafaneHandler)
}
