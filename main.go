package main

import (
    "fmt"
    "github.com/julienschmidt/httprouter"
    "github.com/sirupsen/logrus"
    "log"
    "net/http"
    "os"
)

var logger = logrus.New()
var config *jsonConfig
var wechatWork *WechatWork
var req = Request{}
var version = "0.1"

func main() {
    var router = httprouter.New()
    router.GET("/", index)
    router.POST("/push/", push)
    addr := "0.0.0.0:80"
    logger.Infof("listening at %s", addr)
    log.Fatal(http.ListenAndServe(addr, router))

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
    var err error
    config = loadConfig()
    wechatWork = &WechatWork{config, ""}
    if err != nil {
        logger.Error(err)
    }
}
