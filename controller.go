package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func index(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	fmt.Fprint(res, "post message to /push/")
}

type postMessage struct {
	Message string `json:"message"`
}

func push(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	headerToken := req.Header.Get("token")
	if headerToken != config.Token {
		str := fmt.Sprintf("Incorrect token: %s", headerToken)
		logger.Warnf(str)
		return
	}
	message := req.FormValue("message")
	toUser := req.FormValue("userId")
	err := wechatWork.SendMessage(message, toUser)
	if err != nil {
		fmt.Fprint(res, err.Error())
		return
	}
	fmt.Fprint(res, "OK")
}
