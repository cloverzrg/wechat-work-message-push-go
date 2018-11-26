package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func index(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	fmt.Fprint(res,"post message to /push/")
}

type postMessage struct {
	Message string `json:"message"`
}

func push(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	header_token := req.Header.Get("token")
	if header_token != config.Token{
		str := fmt.Sprintf("Incorrect token: %s",header_token)
		logger.Warnf(str)
		return
	}
	message := req.FormValue("message")

	wechatWork.SendMessage(message)
	fmt.Fprint(res, "OK")
}