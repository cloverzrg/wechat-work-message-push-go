package qyapi

import (
	"encoding/json"
	"fmt"
	"github.com/cloverzrg/wechat-work-message-push-go/config"
	"github.com/cloverzrg/wechat-work-message-push-go/logger"
)

func SendMessage(content string, toUser string) (err error) {
	logger.Infof("push message: %s\n", content)
	token, err := GetToken()
	if err != nil {
		logger.Error(err)
		return err
	}
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%s", token)
	m := TextMessage{
		Agentid: config.Config.WechatWork.AgentId,
		Msgtype: "text",
	}
	if len(toUser) == 0 {
		m.Touser = config.Config.WechatWork.DefaultReceiverUserId
	} else {
		m.Touser = toUser
	}

	m.Text.Content = content

	jsonStr, err := json.Marshal(m)
	if err != nil {
		logger.Error("sendMessage error:%s", err)
		return err
	}
	postJson(url, jsonStr)
	return err
}

func SendCardMessage(content string, title string, imageUrl string, toUser string) (err error) {
	logger.Infof("push message: %s\n", content)
	token, err := GetToken()
	if err != nil {
		return err
	}
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%s", token)
	m := TextCardMessage{
		Msgtype: "textcard",
		Agentid: config.Config.WechatWork.AgentId,
	}
	if len(toUser) == 0 {
		m.Touser = config.Config.WechatWork.DefaultReceiverUserId
	} else {
		m.Touser = toUser
	}

	m.TextCard.Title = title
	m.TextCard.URL = imageUrl
	m.TextCard.Description = content

	jsonStr, err := json.Marshal(m)
	if err != nil {
		logger.Error("sendMessage error:%s", err)
		return err
	}
	postJson(url, jsonStr)
	return err
}