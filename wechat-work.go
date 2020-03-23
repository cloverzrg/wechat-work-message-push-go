package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type WechatWork struct {
	Config    *jsonConfig
	Token     string
	ExpiredAt time.Time
}

func New(config *jsonConfig) *WechatWork {
	return &WechatWork{
		Config: config,
	}
}

type wechatTextMessage struct {
	Touser  string `json:"touser"`
	Msgtype string `json:"msgtype"`
	Agentid string `json:"agentid"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
}

type wechatTextCardMessage struct {
	Touser   string `json:"touser"`
	Msgtype  string `json:"msgtype"`
	Agentid  string `json:"agentid"`
	TextCard struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		URL         string `json:"url"`
	} `json:"textcard"`
}

func (e WechatWork) GetToken() (token string) {
	config := e.Config
	if !e.ExpiredAt.IsZero() && time.Now().Before(e.ExpiredAt) {
		return e.Token
	}
	url := "https://qyapi.weixin.qq.com/cgi-bin/gettoken?" + "corpid=" + config.WechatWork.CorpId + "&corpsecret=" + config.WechatWork.CorpSecret

	data := req.Get(url)
	var objmap map[string]*json.RawMessage
	err := json.Unmarshal(data, &objmap)
	if err != nil {
		logger.Error("getToken error")
		os.Exit(0)
	}

	token = string(*objmap["access_token"])
	token = strings.Replace(token, "\"", "", -1)
	e.Token = token
	logger.Info("new token:", token)
	e.ExpiredAt = time.Now().Add(1 * time.Hour)
	return e.Token
}

func (e WechatWork) SendMessage(content string, toUser string) (err error) {
	logger.Infof("push message: %s\n", content)
	token := e.GetToken()
	config := e.Config
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%s", token)
	m := wechatTextMessage{}
	m.Agentid = config.WechatWork.AgentId
	m.Msgtype = "text"
	if len(toUser) == 0 {
		m.Touser = config.WechatWork.DefaultReceiverUserId
	} else {
		m.Touser = toUser
	}

	m.Text.Content = content

	jsonStr, err := json.Marshal(m)
	if err != nil {
		logger.Error("sendMessage error:%s", err)
		return err
	}
	req.PostJson(url, jsonStr)
	return err
}

func (e WechatWork) SendCardMessage(content string, title string, imageUrl string, toUser string) (err error) {
	logger.Infof("push message: %s\n", content)
	token := e.GetToken()
	config := e.Config
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%s", token)
	m := wechatTextCardMessage{}
	m.Agentid = config.WechatWork.AgentId
	m.Msgtype = "textcard"
	if len(toUser) == 0 {
		m.Touser = config.WechatWork.DefaultReceiverUserId
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
	req.PostJson(url, jsonStr)
	return err
}
