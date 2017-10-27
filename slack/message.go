package slack

type Attachment struct {
	Title string `json:"title"`
	Text  string `json:"text"`
	Color string `json:"color"`
}

type SlackMsg struct {
	Text        string       `json:"text"`
	Username    string       `json:"username"`
	IconUrl     string       `json:"icon_url"`
	Channel     string       `json:"channel"`
	Attachments []Attachment `json:"attachments"`
}

type SlackMsgParams struct {
	Title  string // メッセージのタイトル(ex. Benchmark passed!)
	Text   string // メッセージ本文
	Result bool   // ベンチマークが成功したか否か
	Log    string // エラーログなど
}

func NewSlackMsg(params SlackMsgParams) SlackMsg {
	var color string // カラーコード
	if params.Result {
		// ベンチマーク成功
		color = SUCCESS_COLOR
	} else {
		// ベンチマーク失敗
		color = FAIL_COLOR
	}

	attachments := []Attachment{
		{
			Title: params.Title,
			Text:  params.Text,
			Color: color,
		},
		{
			Title: "- logs - ",
			Text:  params.Log,
		},
	}

	return SlackMsg{
		Text:        "ちっちゃくないもん！",
		Username:    SLACK_USERNAME,
		IconUrl:     SLACK_ICON_URL,
		Channel:     SLACK_CHANNEL,
		Attachments: attachments,
	}
}
