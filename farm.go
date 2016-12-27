package logfarm

import (
	"time"
)

// LogFarm
type LogFarm interface {
	// Set string values
	SetSeparator(string) bool
	// minLength is 102400 bytes = 100 k
	SetMaxLength(int64) bool
	//SetTimerToWriteLog
	SetLoopTimerToWriteLog(time.Duration) bool
	// Write log into cache
	WriteLog(filename string, data []string) bool
}

// New returns logfarm
func New() LogFarm {
	logger := &Logger{
		CurVersion: VerA,
		Separator:  "|",
		Writter:    NewFileWritter(),
	}

	logger.looperWritter()
	return logger
}
