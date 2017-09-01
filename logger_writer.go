// GNU GPL v3 License

// Copyright (c) 2016 github.com:go-trellis

package logfarm

import (
	"github.com/go-trellis/log-farm/proto"
)

const (
	// min length is 102400 bytes
	minLength int64 = 102400
	//
	chanBuffer int = 10000
)

// LoggerWriter logger writter repo
type LoggerWriter interface {
	Write(*logfarm_proto.LogItem) (int, error)
	Stop()
}
