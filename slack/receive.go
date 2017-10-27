package slack

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

func AuthSlackToken(r *http.Request) {
	result, err := ioutil.ReadAll(r.Body)
	if err != nil {
		notifyError(err, "receive.go", "AuthSlackToken", "HTTPリクエストのRead中にエラーが起きました.")
		panic(err)
	}

	params, err := url.ParseQuery(string(result))
	if err != nil {
		notifyError(err, "receive.go", "AuthSlackToken", "HTTPリクエストのParseQuery中にエラーがh起きました.")
		panic(err)
	}

	token := params["token"][0]
	if token != OUTGOING_WEBHOOK_TOKEN {
		notifyError(err, "receive.go", "AuthSlackToken", "不正なトークンを受信しました.")
		panic(err)
	}
}
