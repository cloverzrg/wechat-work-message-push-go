package qyapi

import (
	"encoding/json"
	"github.com/cloverzrg/wechat-work-message-push-go/config"
	"github.com/cloverzrg/wechat-work-message-push-go/logger"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type tokenSer struct {
	Token     string
	ExpiredAt time.Time
}

var token tokenSer

func GetToken() (tokenStr string, err error) {
	if !token.ExpiredAt.IsZero() && time.Now().Before(token.ExpiredAt) {
		return token.Token, err
	}
	url := "https://qyapi.weixin.qq.com/cgi-bin/gettoken?" + "corpid=" + config.Config.WechatWork.CorpId + "&corpsecret=" + config.Config.WechatWork.CorpSecret

	resp, err := http.Get(url)
	if err != nil {
		logger.Error(err)
		return tokenStr, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	var objmap map[string]*json.RawMessage
	err = json.Unmarshal(data, &objmap)
	if err != nil {
		logger.Error("getToken error")
		return tokenStr, err
	}

	tokenStr = string(*objmap["access_token"])
	tokenStr = strings.Replace(tokenStr, "\"", "", -1)
	token.Token = tokenStr
	logger.Info("new token:", token)
	token.ExpiredAt = time.Now().Add(1 * time.Hour)
	return token.Token, err
}
