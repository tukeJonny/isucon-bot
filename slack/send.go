package slack

func SendSlack(w http.ResponseWriter, message string) {
	msg, err := json.Marshal(NewSlackMsg(SlackMsgParams{
		Title: "Notification",
		Text: message,
		Result: true,
		Log: "",
	})
	if err != nil {
		notifyError(err, "main.go", "SendSlack", "JSONのMarshal中にエラーが起きました")
		panic(err)
	}
	fmt.Fprintf(w, string(msg))
}


