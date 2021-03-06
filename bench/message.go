package bench

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/tukejonny/isucon-bot/slack"
)

type BenchResult struct {
	JobID     string    `json:"job_id"`
	Score     int64     `json:"score"`
	Pass      bool      `json:"pass"` // 成功したか否か
	Message   string    `json:"message"`
	LoadLevel int64     `json:"load_level"`
	IpAddrs   string    `json:"ip_addrs"` // 複数指定できるけど、ぼっと側での判断が面倒なので
	Error     []string  `json:"error"`
	Log       []string  `json:"log"`
	StartAt   time.Time `json:"start_time"`
	EndAt     time.Time `json:"end_time"`
}

func (result *BenchResult) GetSlackMsg() slack.SlackMsg {
	var title string
	if result.Pass {
		title = "Benchmark passed. :heart:"
	} else {
		title = "Benchmark failed. :broken_heart:"
	}

	msg := ""
	msg = msg + fmt.Sprintf("[Benchmark %s ~ %s]", result.StartAt.String(), result.EndAt.String()) + "\n"
	msg = msg + fmt.Sprintf("[JobID for %s]: ", result.IpAddrs) + result.JobID
	msg = msg + "[LoadLevel]: " + strconv.FormatInt(result.LoadLevel, 10) + "\n"
	msg = msg + "[Score]: " + strconv.FormatInt(result.Score, 10) + "\n"
	msg = msg + "[Message]: " + result.Message + "\n"

	log := ""
	log = log + "[ERROR LOG]\n"
	log = log + strings.Join(result.Error[:logLimit], "\n")
	log = log + "\n[LOG]\n"
	log = log + strings.Join(result.Log[:logLimit], "\n")

	return slack.NewSlackMsg(slack.SlackMsgParams{
		Title:  title,
		Text:   msg,
		Result: result.Pass,
		Log:    log,
	})
}
