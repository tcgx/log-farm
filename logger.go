package logfarm

import (
	"sync"
	"time"
)

type Logger struct {
	CurVersion string
	Separator  string

	Writter LoggerWritter

	timer time.Duration

	sync.RWMutex
}

func NewLogger() LogFarm {
	logger := &Logger{
		CurVersion: VerA,
		Separator:  "|",
	}

	return logger
}

type LogItem struct {
	CreateTime time.Time
	Filename   string
	Values     []string
	Separator  string
}

func (p *Logger) SetSeparator(s string) bool {
	if s == "" {
		return false
	}
	p.Lock()
	defer p.Unlock()

	p.Separator = s
	return p.Separator == s
}

func (p *Logger) SetMaxFileLength(l int64) bool {
	return p.Writter.SetMaxFileLength(l)
}

func (p *Logger) WriteLog(filename string, data []string) bool {
	if filename == "" {
		return false
	}
	item := &LogItem{
		CreateTime: time.Now(),
		Filename:   filename,
		Values:     data,
		Separator:  p.Separator,
	}
	return Cache.Insert(p.CurVersion, item.Filename, item)
}

func (p *Logger) SetTimerToWriteLog(t time.Duration) bool {
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
