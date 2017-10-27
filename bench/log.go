package bench

import (
	log "github.com/sirupsen/logrus"
)

func WriteBenchLog(params BenchResult) {
	var msg string
	if params.Pass {
		msg = "benchmark passed"
	} else {
		msg = "benchmark failed"
	}
	log.WithFields(log.Fields{
		"job_id":    params.JobID,
		"score":     params.Score,
		"message":   params.Message,
		"loadlevel": params.LoadLevel,
		"ipaddr":    params.IpAddrs,
		"start_at":  params.StartAt,
		"end_at":    params.EndAt,
	}).Info(msg)
}
