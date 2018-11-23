package main

import (
	"github.com/sirupsen/logrus"
	"os"
)

var logger = logrus.New()
var req = Request{}
var version = "0.1"

func main() {
	params := parseCmdParams()
	if params.isPrintVersion {
		logger.Infof("version: %s", version)
		os.Exit(0)
	}
	config, err := loadConfig(params.configPath)
	wechatWork := WechatWork{config,""}
	if err != nil {
		print(err)
	}
	wechatWork.SendMessage("213")

}
