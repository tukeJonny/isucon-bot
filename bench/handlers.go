package bench

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/tukejonny/isucon-bot/slack"
)

var (
	BenchmarkLock bool // ベンチマークロックが確保可能であるか判別するためにsync.Mutexを用いない

	basePath   = "/home/isucon/isubata/bench"
	resultPath = "/home/isucon/isubata/bench/result.json"
)

func benchmark(target string) {
	// Execute benchmark
	os.Chdir(basePath)
	_, err := exec.Command("./bin/bench", "-remotes", target, "-output", "result.json").Output()
	if err != nil {
		slack.NotifyErr(err, "handlers.go", "benchmark", "ベンチマーク実行中にエラーが起きました")
		panic(err)
	}

	// Get Result of benchmark
	resultJson, err := ioutil.ReadFile(resultPath)
	if err != nil {
		slack.NotifyErr(err, "handlers.go", "benchmark", "result.jsonを読む途中でエラーが起きました")
		panic(err)
	}

	// Notify Result to slack
	var result BenchResult
	err = json.Unmarshal([]byte(resultJson), &result)
	if err != nil {
		slack.NotifyErr(err, "handlers.go", "benchmark", "result.jsonのUnmarshalize中でエラーが起きました")
		panic(err)
	}

	// Make message
	slackMsg := result.GetSlackMsg()
	message, err := json.Marshal(&slackMsg)
	if err != nil {
		slack.NotifyErr(err, "handlers.go", "benchmark", "SlackMsgをMarshalizeする途中でエラーが起きました")
		panic(err)
	}

	slack.SendSlack(string(message))
}

func BenchmarkHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if BenchmarkLock {
		slack.RespondSlack(w, "まだ実行中です♪ もうちょっと待ってね!")
	} else {
		params := slack.GetSlackParams(r)
		slack.AuthSlackToken(params)

		userParameters := strings.Split(params["text"][0], " ")
		target := benchTargets[userParameters[1]]

		go func() {
			BenchmarkLock = true
			benchmark(target)
			BenchmarkLock = false
		}()
		slack.RespondSlack(w, "ベンチマークを走らせています♪")
	}
}
