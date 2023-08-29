package qyapi

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	logger.Infof("post response:%s", body)
	var objmap map[string]*json.RawMessage
	err = json.Unmarshal(body, &objmap)
	if err != nil {
		logger.Error("postJson error")
		os.Exit(0)
	}

	errcode := string(*objmap["errcode"])
	if errcode != "0" {
		err = fmt.Errorf("postJson errmsg:" + string(*objmap["errmsg"]))
		logger.Errorf(err.Error())
		return body, err
	}
	return body, err
}
