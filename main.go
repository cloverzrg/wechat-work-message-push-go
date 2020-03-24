package main

import (
	"fmt"
	"github.com/cloverzrg/wechat-work-message-push-go/api"
	"github.com/cloverzrg/wechat-work-message-push-go/config"
	"github.com/cloverzrg/wechat-work-message-push-go/logger"
	"os"
)

var (
	BuildTime string
	GoVersion string
	GitHead   string
)

func main() {
	err := api.Start()
	logger.Error(err)
}

func init() {
	buildInfo := fmt.Sprintf("BuildTime: %s\nGoVersion: %s\nGitHead: %s\n", BuildTime, GoVersion, GitHead)
	if BuildTime != "" {
		fmt.Print(buildInfo)
	}
	params := config.ParseCmdParams()
	if params.IsPrintVersion {
		os.Exit(0)
	}
	err := config.LoadConfig()
	if err != nil {
		logger.Error(err)
		return
	}
}
