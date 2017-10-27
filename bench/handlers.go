package bench

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"

	"github.com/julienschmidt/httprouter"
	"github.com/tukejonny/isucon-bot/slack"
)

var (
	BenchmarkLock bool // ベンチマークロックが確保可能であるか判別するためにsync.Mutexを用いない

	basePath   = "/home/isucon/isubata/bench"
	resultPath = "/home/isucon/isubata/bench/result.json"
)

func benchmark() {
	// Execute benchmark
	os.Chdir(basePath)
	_, err := exec.Command("./bin/bench", "").Output()
	if err != nil {
		notifyErr(err, "handlers.go", "benchmark", "ベンチマーク実行中にエラーが起きました")
		panic(err)
	}

	// Get Result of benchmark
	result, err := ioutil.ReadFile(resultPath)
	if err != nil {
		notifyErr(err, "handlers.go", "benchmark", "result.jsonを読む途中でエラーが起きました")
		panic(err)
	}

	// Notify Result to slack
	var result BenchResult
	err = json.Unmarshal([]byte(message), &result)
	if err != nil {
		notifyErr(err, "handlers.go", "benchmark", "result.jsonのUnmarshalize中でエラーが起きました")
		panic(err)
	}

	slack.SendSlack(result.GetSlackMsg())
}

func BenchmarkHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if benchmarkLock {
		respond(w, "まだ実行中です♪ もうちょっと待ってね!")
	} else {
		AuthSlackToken(r)
		go func() {
			BenchmarkLock = true
			benchmark()
			BenchmarkLock = false
		}()
		respond(w, "ベンチマークを走らせています♪")
	}
}
