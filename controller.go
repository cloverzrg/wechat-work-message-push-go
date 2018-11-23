package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func index(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	ua := req.UserAgent()
	fmt.Fprint(res, "Hello World!"+"\n")
	fmt.Fprint(res, id)
	fmt.Fprint(res, ua)
}

type postMessage struct {
	Message string
}

func push(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	header_token := req.Header.Get("token")
	if header_token != config.Token{
		str := fmt.Sprintf("Incorrect token: %s",header_token)
		logger.Warnf(str)
		fmt.Fprint(res, str)
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