package logger

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

func GinLogger() gin.HandlerFunc {
	log := Entry
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		end := time.Now()
		latency := end.Sub(start)
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()

		entry := log.WithFields(logrus.Fields{
			"statusCode": statusCode,
			"latency":    latency,
			"clientIP":   clientIP,
			"method":     c.Request.Method,
			"loggerType": "gin",
			"requestURI": c.Request.RequestURI,
		})

		msg := ""
		if statusCode != 200 {
			entry.Errorf(msg)
		} else {
			entry.Info(msg)
		}
	}
}
