package model

import "time"

type LoggerReq struct {
	StartTime  time.Time
	EndTime    time.Duration
	StatusCode int
	ClientIP   string
	Method     string
	Path       string
	UserAgent  string
}
