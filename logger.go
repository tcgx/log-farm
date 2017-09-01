// GNU GPL v3 License

// Copyright (c) 2016 github.com:go-trellis

package logfarm

import (
	"github.com/go-trellis/log-farm/proto"
)

// logger implements for logger writer
type logger struct {
	Writer LoggerWriter

	logChan  chan *logfarm_proto.LogItem
	stopChan chan bool
}

// WriteLog write logs to the filename
func (p *logger) WriteLog(data []string) bool {

	p.logChan <- &logfarm_proto.LogItem{Values: data}

	return true
}

// Stop stop write data into file
func (p *logger) Stop() {
	p.stopChan <- true
}

func (p *logger) looperWriter() {
	go func() {
		for {
			select {
			case log := <-p.logChan:
				p.Writer.Write(log)
			case <-p.stopChan:
				p.Writer.Stop()
				close(p.logChan)
				close(p.stopChan)
				return
			}
		}
	}()
}
