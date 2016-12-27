package logfarm

import (
	"time"
)

type LogFarm interface {
	// Set string values
	SetSeparator(string) bool
	// minFileLength is 1024 bytes
	SetMaxFileLength(int64) bool
	//SetTimerToWriteLog
	SetTimerToWriteLog(time.Duration) bool
	// Write log into cache
	WriteLog(filename string, data []string) bool
}
