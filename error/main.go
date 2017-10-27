
func notifyError(err error, file string, function string, message string) {
	var msg string = fmt.Sprintf("%s:%s: %s.", file, function, message)
	json := fmt.Sprintf("{'text':'[!] %v\n'}", errors.Wrap(err, msg))
	http.PostForm(
		INCOMING_WEBHOOK_URL,
		url.Values{"payload": {json}},
	)
}
