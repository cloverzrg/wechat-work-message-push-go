package main

import (
    "fmt"
    "github.com/julienschmidt/httprouter"
    "github.com/sirupsen/logrus"
    "net/http"
    "os"
    "time"
)

var logger = logrus.New()
var config *jsonConfig
var wechatWork *WechatWork
var req = Request{}
var version = "0.3"

func main() {
	var router = httprouter.New()
	router.GET("/", index)
	router.POST("/push/", push)
	router.POST("/grafana/", GrafaneHandler)
	addr := "0.0.0.0:80"
	logger.Infof("listening at %s", addr)
    logger.Fatal(http.ListenAndServe(addr, router))
}

func init() {
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	params := parseCmdParams()
	if params.isPrintVersion {
		fmt.Printf("version: %s", version)
		os.Exit(0)
	}
	config = loadConfig()
	wechatWork = &WechatWork{
		Config:    config,
		Token:     "",
		ExpiredAt: time.Time{},
	}
}
