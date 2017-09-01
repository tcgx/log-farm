// GNU GPL v3 License

// Copyright (c) 2016 github.com:go-trellis

package logfarm

import (
	"sync"

	"github.com/go-trellis/config"
	"github.com/go-trellis/log-farm/proto"
)

const (
	namespace = "Trellis::LogFarm"
)

// LogFarm functions to wite logs
type LogFarm interface {
	// Write log into cache
	WriteLog(data []string) bool
	// stop write data into file
	Stop()
}

var mapLogger = make(map[string]*logger)
var mapLoggerLocker sync.Mutex

// New returns logfarm
// Initial params
// filename: the log path & name, see more in example
// filesuffix: log file' suffix
// chanbuffer: length of the log chan buffer
// filemaxlength: the max length of log file, default: 0 is unlimit
// movefiletype: move file by per-minite(1) or hourly(2) or daily(3), 0 is doing nothing
func New(filename string, options config.Options) LogFarm {
	mapLoggerLocker.Lock()
	log := mapLogger[filename]
	if log != nil {
		mapLoggerLocker.Unlock()
		return log
	}
	log = &logger{
		Writer: NewFileWriter(filename, options),
	}

	_chanBuffer, err := options.Int("chanbuffer")
	if err != nil {
		panic(err)
	}
	log.logChan = make(chan *logfarm_proto.LogItem, _chanBuffer)
	log.stopChan = make(chan bool)
	log.looperWriter()

	mapLogger[filename] = log
	mapLoggerLocker.Unlock()
	return log
}
