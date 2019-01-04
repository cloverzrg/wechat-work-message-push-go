package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type WechatWork struct {
	Config *jsonConfig
	Token string
	//TokenExpired int
}

func New(config *jsonConfig) *WechatWork {
	return &WechatWork{
		Config: config,
	}
}

type wechatMessage struct {
	Touser  string `json:"touser"`
	Msgtype string `json:"msgtype"`
	Agentid string `json:"agentid"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
}

func (e WechatWork)  GetToken() (token string) {
	config := e.Config
	url := "https://qyapi.weixin.qq.com/cgi-bin/gettoken?" + "corpid=" + config.WechatWork.CorpId + "&corpsecret=" + config.WechatWork.CorpSecret

	data := req.Get(url)
	var objmap map[string]*json.RawMessage
	err := json.Unmarshal(data, &objmap)
	if err != nil {
		logger.Infof("getToken error")
		os.Exit(0)
	}

	token = string(*objmap["access_token"])
	token = strings.Replace(token,"\"","",-1)
	e.Token = token
	return e.Token
}

func (e WechatWork)  SendMessage(content string, toUser string) bool {
	logger.Infof("push message: %s\n",content)
	token := e.GetToken()
	config := e.Config
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%s", token)
	m := wechatMessage{}
	m.Agentid = config.WechatWork.AgentId
	m.Msgtype = "text"
	if len(toUser)==0 {
		m.Touser = config.WechatWork.DefaultReceiverUserId
	}else{
		m.Touser = toUser
	}

	m.Text.Content = content

	jsonStr,err := json.Marshal(m)
	if err != nil {
		logger.Infof("sendMessage error:%s",err)
	}
	req.PostJson(url, jsonStr)
	return true
}
