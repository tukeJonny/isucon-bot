package slack

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

func GetSlackParams(r *http.Request) url.Values {
	result, err := ioutil.ReadAll(r.Body)
	if err != nil {
		NotifyErr(err, "receive.go", "GetSlackParams", "HTTPリクエストのRead中にエラーが起きました")
		panic(err)
	}

	params, err := url.ParseQuery(string(result))
	if err != nil {
		NotifyErr(err, "receive.go", "GetSlackParams", "HTTPリクエストのParseQuery中にエラーが起きました")
	}

	return params
}

func AuthSlackToken(params url.Values) {
	token := params["token"][0]
	if token != OUTGOING_WEBHOOK_TOKEN {
		err := errors.New("Invalid OUTGOING TOKEN")
		NotifyErr(err, "receive.go", "AuthSlackToken", "不正なトークンを受信しました.")
		panic(err)
	}
}
