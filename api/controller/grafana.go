package controller

import (
	"encoding/json"
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
	noti := grafana.Notification{}
	err = json.Unmarshal(bytes, &noti)
	//err = c.BindJSON(&noti)
	if err != nil {
		c.String(500, err.Error())
		return
	}
	err = qyapi.SendCardMessage(noti.Message, noti.Title, noti.ImageUrl, "")
	if err != nil {
		c.String(500, err.Error())
		return
	}
	c.String(200, "ok")
}
