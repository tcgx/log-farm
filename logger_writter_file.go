package logfarm

import (
	"strings"
	"sync"
	"time"

	"github.com/go-trellis/files"
	"github.com/go-trellis/formats/times"
	"github.com/go-trellis/log-farm/proto"
)

type fileWritter struct {
	sync.RWMutex

	MaxLength int64
}

var fileExecutor = files.New()

func NewFileWritter() LoggerWritter {
	return &fileWritter{
		MaxLength: minLength,
	}
}

func (p *fileWritter) SetMaxLength(l int64) bool {
	p.Lock()
	defer p.Unlock()
	if l > minLength {
		p.MaxLength = l
	}
	return p.MaxLength == l
}

func (p *fileWritter) Write(tab string) (n int64, err error) {
	p.Lock()
	defer p.Unlock()

	keys, ok := Cache.Members(tab)
	if !ok {
		return 0, nil
	}

	var size int64

	for _, name := range keys {
		if fi, _ := fileExecutor.FileInfo(name); fi == nil {
			size = 0
		} else {
			size = fi.Size()
		}

		values, ok := Cache.Lookup(tab, name)
		if !ok {
			continue
		}

		for _, v := range values {
			log, ok := v.(logfarm_proto.LogItem)
			if !ok {
				continue
			}

			count, e := fileExecutor.WriteAppend(name, strings.Join(
				append([]string{log.CreateTime}, log.Values...), log.Separator)+
				"\n")
			if e != nil {
				err = e
				return
			}
			n += int64(count)
			size += int64(count)
			if size < p.MaxLength {
				continue
			}

			if err = p.moveFile(name); err != nil {
				return
			}
			size = 0
		}
	}
	return
}

func (p *fileWritter) ResetTab(tab string) bool {
	p.Lock()
	defer p.Unlock()
	return Cache.DeleteAllObjects(tab)
}

func (p *fileWritter) moveFile(name string) error {
	return fileExecutor.Rename(name, name+"."+times.TimeToRFC3339Nano(time.Now()))
}
