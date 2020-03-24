package qyapi

type TextMessage struct {
	Touser  string `json:"touser"`
	Msgtype string `json:"msgtype"`
	Agentid string `json:"agentid"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
}

type TextCardMessage struct {
	Touser   string `json:"touser"`
	Msgtype  string `json:"msgtype"`
	Agentid  string `json:"agentid"`
	TextCard struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		URL         string `json:"url"`
	} `json:"textcard"`
}
