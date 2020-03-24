package config

import (
	"fmt"
	"os"
)

type jsonConfig struct {
	Token      string
	WechatWork struct {
		DefaultReceiverUserId string
		CorpId                string
		CorpSecret            string
		AgentId               string
	}
	GrafanaWebhookUser     string
	GrafanaWebhookPassword string
}

var Config jsonConfig

func LoadConfig() (err error) {

	Config.Token = os.Getenv("Token")
	Config.WechatWork.CorpSecret = os.Getenv("WechatWorkCorpSecret")
	Config.WechatWork.CorpId = os.Getenv("WechatWorkCorpId")
	Config.WechatWork.DefaultReceiverUserId = os.Getenv("DefaultReceiverUserId")
	Config.WechatWork.AgentId = os.Getenv("WechatWorkAgentId")
	Config.GrafanaWebhookUser = os.Getenv("GrafanaWebhookUser")
	Config.GrafanaWebhookPassword = os.Getenv("GrafanaWebhookPassword")
	fmt.Printf("%+v\n", Config)
	return err
}
