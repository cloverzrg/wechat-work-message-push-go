package main

import (
	"crypto/subtle"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
)

type Notification struct {
	DashboardId int `json:"dashboardId"`
	EvalMatches []struct {
		Value  int         `json:"value"`
		Metric string      `json:"metric"`
		Tags   interface{} `json:"tags"`
	} `json:"evalMatches"`
	ImageUrl string      `json:"imageUrl"`
	Message  string      `json:"message"`
	OrgId    int         `json:"orgId"`
	PanelId  int         `json:"panelId"`
	RuleId   int         `json:"ruleId"`
	RuleName string      `json:"ruleName"`
	RuleUrl  string      `json:"ruleUrl"`
	State    string      `json:"state"`
	Tags     interface{} `json:"tags"`
	Title    string      `json:"title"`
}

func GrafaneHandler(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	user, pass, ok := req.BasicAuth()

	if !ok || subtle.ConstantTimeCompare([]byte(user), []byte(config.GrafanaWebhookUser)) != 1 || subtle.ConstantTimeCompare([]byte(pass), []byte(config.GrafanaWebhookPassword)) != 1 {
		res.Header().Set("WWW-Authenticate", `Basic realm="`+"Unauthorized"+`"`)
		res.WriteHeader(401)
		res.Write([]byte("Unauthorized.\n"))
		return
	}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Fprint(res, err.Error())
		return
	}
	noti := Notification{}
	err = json.Unmarshal(body, &noti)
	if err != nil {
		fmt.Fprint(res, err.Error())
		return
	}
	err = wechatWork.SendCardMessage(noti.Message, noti.Title, noti.ImageUrl, "")
	if err != nil {
		fmt.Fprint(res, err.Error())
		return
	}
	fmt.Fprint(res, "OK")
}
