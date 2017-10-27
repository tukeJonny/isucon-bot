package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tukejonny/isucon-bot/bench"
	"github.com/tukejonny/isucon-bot/slack"
)

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
