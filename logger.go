// GNU GPL v3 License

// Copyright (c) 2016 github.com:go-trellis

package logfarm

import (
	"sync"
	"time"

	"github.com/go-trellis/formats"
	"github.com/go-trellis/log-farm/proto"
)

// Logger implements for logger writter
type Logger struct {
	CurVersion string
	Separator  string

	Writter LoggerWritter

	timer time.Duration

	sync.RWMutex
}

// SetSeparator set Separator for a log's colunms
func (p *Logger) SetSeparator(s string) bool {
	if s == "" {
		return false
	}
	p.Lock()
	defer p.Unlock()

	p.Separator = s
	return p.Separator == s
}

// SetMaxLength set max length
func (p *Logger) SetMaxLength(l int64) bool {
	return p.Writter.SetMaxLength(l)
}

// WriteLog write logs to the filename
func (p *Logger) WriteLog(filename string, data []string) bool {
	p.Lock()
	defer p.Unlock()
	if filename == "" {
		return false
	}
	item := logfarm_proto.LogItem{
		CreateTime: formats.FormatDashTime(time.Now()),
		Filename:   filename,
		Values:     data,
		Separator:  p.Separator,
	}

	return Cache.Insert(p.CurVersion, item.Filename, item)
}

// SetLoopTimerToWriteLog looper for writting logs
func (p *Logger) SetLoopTimerToWriteLog(t time.Duration) bool {
	p.Lock()
	defer p.Unlock()
	p.timer = t
	return p.timer == t
}

func (p *Logger) changeVer() {
	p.Lock()
	defer p.Unlock()

	switch p.CurVersion {
	case VerA:
		p.CurVersion = VerB
	case VerB:
		p.CurVersion = VerC
	case VerC:
		p.CurVersion = VerA
	default:
		p.CurVersion = VerA
	}
}

func (p *Logger) getBackVer() string {
	switch p.CurVersion {
	case VerA:
		return VerC
	case VerB:
		return VerA
	case VerC:
		return VerB
	default:
		return VerB
	}
}

func (p *Logger) looperWritter() {
	go func() {
		for {
			p.changeVer()

			if _, e := p.write(p.getBackVer()); e != nil {
				// TODO Log
				continue
			}
			time.Sleep(p.timer)
		}
	}()
}

func (p *Logger) write(ver string) (n int64, err error) {

	if n, err = p.Writter.Write(ver); err != nil {
		return
	}
	p.Writter.ResetTab(ver)
	return
}
