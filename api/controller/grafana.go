package controller

import (
	"encoding/json"
	"fmt"
	"github.com/cloverzrg/wechat-work-message-push-go/grafana"
	"github.com/cloverzrg/wechat-work-message-push-go/logger"
	"github.com/cloverzrg/wechat-work-message-push-go/qyapi"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func GrafaneHandler(c *gin.Context) {
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info("GrafaneHandler: ", string(bytes))
	notification := grafana.Notification{}
	err = json.Unmarshal(bytes, &notification)
	//err = c.BindJSON(&notification)
	if err != nil {
		c.String(500, err.Error())
		return
	}
	msg := notification.Title
	if notification.Message != "" {
		msg += "\n" + notification.Message
	}
	if len(notification.EvalMatches) > 0 {
		msg += "\nAlert matched metrics:"
		for _, v := range notification.EvalMatches {
			msg += fmt.Sprintf("\n%s: %f", v.Metric, v.Value)
		}
	}
	//err = qyapi.SendCardMessage(notification.Message, notification.Title, notification.ImageUrl, "")
	err = qyapi.SendMessage(msg, "", "grafana")
	if err != nil {
		c.String(500, err.Error())
		return
	}
	c.String(200, "ok")
}
