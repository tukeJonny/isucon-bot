package main

import (
	"bufio"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	"github.com/tukejonny/isucon-bot/bench"
	"github.com/tukejonny/isucon-bot/slack"
)

const (
	LOG_PATH = "./bench.log"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})

	fd, err := os.OpenFile(LOG_PATH, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		slack.NotifyErr(err, "logger:main.go", "init", "ログファイルを開けませんでした")
		panic(err)
	}

	writer := bufio.NewWriter(fd)
	log.SetOutput(writer)

	log.SetLevel(log.DebugLevel)
}

func initRouter() (router *httprouter.Router) {
	router = httprouter.New()
	router.POST("/bench", bench.BenchmarkHandler)

	return
}

func main() {
	router := initRouter()
	if err := http.ListenAndServe(":8080", router); err != nil {
		slack.NotifyErr(err, "main.go", "main", "http.ListenAndServe(port=8080)中にエラーが発生しました.")
	}
}
