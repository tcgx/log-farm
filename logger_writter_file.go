// GNU GPL v3 License

// Copyright (c) 2016 github.com:go-trellis

package logfarm

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/go-trellis/config"
	"github.com/go-trellis/files"
	"github.com/go-trellis/log-farm/proto"
)

// MoveFileType move file type
type MoveFileType int

// MoveFileTypes
const (
	MoveFileTypeNothing MoveFileType = iota
	MoveFileTypePerMinite
	MoveFileTypeHourly
	MoveFileTypeDaily
)

type fileWritter struct {
	locker sync.Mutex

	MaxLength int64

	writeFile  string
	FileName   string
	FileSuffix string
	Separator  string

	TimeToMove   MoveFileType
	ticker       *time.Ticker
	lastMoveFlag int
}

var fileExecutor = files.New()

// NewFileWritter get file logger writter
func NewFileWritter(filename string, options config.Options) LoggerWritter {
	fw := &fileWritter{
		FileName:  filename,
		Separator: "|",
	}
	var err error
	if 0 == len(fw.FileName) {
		panic("filename not exist, input with param filename")
	}
	fw.writeFile = fw.FileName

	if options == nil {
		return fw
	}

	fw.MaxLength, err = options.Int64("filemaxlength")
	if err != nil {
		panic(err)
	}

	fw.FileSuffix, err = options.String("filesuffix")
	if err != nil {
		panic(err)
	} else if 0 != len(fw.FileSuffix) {
		fw.writeFile += "." + fw.FileSuffix
	}

	_separator, err := options.String("separator")
	if err != nil {
		panic(err)
	} else if 0 != len(_separator) {
		fw.Separator = _separator
	}

	_movefiletype, err := options.Int("movefiletype")
	if err != nil {
		panic(err)
	}
	fw.TimeToMove = MoveFileType(_movefiletype)
	switch fw.TimeToMove {
	case MoveFileTypePerMinite:
		fw.lastMoveFlag = time.Now().Minute()
		fallthrough
	case MoveFileTypeHourly:
		fw.lastMoveFlag = time.Now().Hour()
		fallthrough
	case MoveFileTypeDaily:
		fw.lastMoveFlag = time.Now().Day()
		fw.ticker = time.NewTicker(time.Second)
		fw.timeToMoveFile()
	}

	return fw
}

func (p *fileWritter) Write(log *logfarm_proto.LogItem) (n int, err error) {

	p.locker.Lock()
	defer p.locker.Unlock()

	p.judgeMoveFile()

	n, err = fileExecutor.WriteAppend(p.writeFile, strings.Join(log.Values, p.Separator)+"\n")
	if err != nil {
		return
	}

	if p.MaxLength == 0 {
		return
	}

	fi, e := fileExecutor.FileInfo(p.writeFile)
	if e != nil {
		return 0, e
	}

	if p.MaxLength > fi.Size() {
		return
	}

	p.moveFile(time.Now().Format("20060102T150405.999999999"))

	return
}

func (p *fileWritter) judgeMoveFile() error {

	timeStr, flag := "", 0
	timeNow := time.Now()
	switch p.TimeToMove {
	case MoveFileTypePerMinite:
		flag = timeNow.Minute()
		timeStr = fmt.Sprintf("%s", timeNow.Format("20060102150405"))
	case MoveFileTypeHourly:
		flag = timeNow.Hour()
		timeStr = fmt.Sprintf("%s%0.2d", timeNow.Format("20060102"), flag)
	case MoveFileTypeDaily:
		flag = time.Now().Day()
		timeStr = timeNow.Format("20060102")
	default:
		return nil
	}

	if flag == p.lastMoveFlag {
		return nil
	}
	p.lastMoveFlag = flag
	return p.moveFile(timeStr)
}

func (p *fileWritter) moveFile(timeStr string) error {
	filename := fmt.Sprintf("%s_%s", p.FileName, timeStr)
	if 0 != len(p.FileSuffix) {
		filename += "." + p.FileSuffix
	}
	return fileExecutor.Rename(p.writeFile, filename)
}

func (p *fileWritter) timeToMoveFile() {
	go func() {
		for {
			select {
			case t := <-p.ticker.C:
				flag := 0
				switch p.TimeToMove {
				case MoveFileTypePerMinite:
					flag = t.Minute()
				case MoveFileTypeHourly:
					flag = t.Hour()
				case MoveFileTypeDaily:
					flag = t.Day()
				}
				if p.lastMoveFlag == flag {
					continue
				}
				p.locker.Lock()
				p.judgeMoveFile()
				p.locker.Unlock()
			}
		}
	}()
}
