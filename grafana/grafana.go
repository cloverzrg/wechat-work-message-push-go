package grafana

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
