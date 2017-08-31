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
func New(filename string, options config.Options) LogFarm {
	mapLoggerLocker.Lock()
	log := mapLogger[filename]
	if log != nil {
		mapLoggerLocker.Unlock()
		return log
	}
	log = &logger{
		Writter: NewFileWritter(filename, options),
	}

	_chanBuffer, err := options.Int("chan_buffer")
	if err != nil {
		panic(err)
	}
	log.logChan = make(chan *logfarm_proto.LogItem, _chanBuffer)
	log.stopChan = make(chan bool)
	log.looperWritter()

	mapLogger[filename] = log
	mapLoggerLocker.Unlock()
	return log
}
