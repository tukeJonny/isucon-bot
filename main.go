package main

func initRouter() (router *httpRouter.Router) {
	router = httprouter.New()
	router.POST("/bench", benchHandler)

	return
}

func main() {
	router := initRouter()
	if err := http.ListenAndServe(":8080", router); err != nil {
		notifyError(err, "main.go", "main", "http.ListenAndServe(port=8080)中にエラーが発生しました.")
	}
}
