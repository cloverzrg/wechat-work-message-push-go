package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func index(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	fmt.Fprint(res,"post message to /push/")
}

type postMessage struct {
	Message string
}

func push(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	header_token := req.Header.Get("token")
	if header_token != config.Token{
		str := fmt.Sprintf("Incorrect token: %s",header_token)
		logger.Warnf(str)
		return
	}
	decoder := json.NewDecoder(req.Body)
	var m postMessage
	err := decoder.Decode(&m)
	if err != nil {
		panic(err)
	}

	wechatWork.SendMessage(m.Message)
	fmt.Fprint(res, "OK")
}