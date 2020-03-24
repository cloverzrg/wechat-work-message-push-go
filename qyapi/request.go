package qyapi

import (
	"bytes"
	"encoding/json"
	"github.com/cloverzrg/wechat-work-message-push-go/logger"
	"io/ioutil"
	"net/http"
	"os"
)

func postJson(url string, jsonStr []byte) (body []byte, err error) {

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ = ioutil.ReadAll(resp.Body)
	var objmap map[string]*json.RawMessage
	err = json.Unmarshal(body, &objmap)
	if err != nil {
		logger.Error("postJson error")
		os.Exit(0)
	}

	errcode := string(*objmap["errcode"])
	if errcode != "0" {
		logger.Errorf("postJson errmsg:" + string(*objmap["errmsg"]))
	}
	return body, err
}
