package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type jsonConfig struct {
	Host       string
	Port       int
	Token      string
	WechatWork struct {
		ReceiverUserId string
		Corpid         string
		Corpsecret     string
		Agentid        string
	}
}

func loadConfig(path string) (config *jsonConfig, err error) {
	logger.Infof("reading config from %s", path)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)

	if err != nil {
		return nil, err
	}
	config = &jsonConfig{}
	if err = json.Unmarshal(data, config); err != nil {
		return nil, err
	}
	return
}
