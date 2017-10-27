package slack

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

func NotifyErr(err error, file string, function string, message string) {
	var msg string = fmt.Sprintf("%s:%s: %s.", file, function, message)
	json := fmt.Sprintf("{'text':'[!] %v\n'}", errors.Wrap(err, msg))
	SendSlack(json)
}

func SendSlack(message string) {
	http.PostForm(
		INCOMING_WEBHOOK_URL,
		url.Values{"payload": {message}},
	)
}

func RespondSlack(w http.ResponseWriter, message string) {
	msg, err := json.Marshal(NewSlackMsg(SlackMsgParams{
		Title:  "Notification",
		Text:   message,
		Result: true,
		Log:    "",
	}))
	if err != nil {
		NotifyErr(err, "main.go", "SendSlack", "JSONのMarshal中にエラーが起きました")
		panic(err)
	}
	fmt.Fprintf(w, string(msg))
}
