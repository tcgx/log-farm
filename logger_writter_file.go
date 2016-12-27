package logfarm

import (
	"strings"
	"sync"

	"github.com/go-trellis/files"
	"github.com/go-trellis/formats/times"
)

type fileWitter struct {
	sync.RWMutex

	MaxFileLength int64
}

var (
	fileExecutor = files.New()
)

func (p *fileWitter) SetMaxFileLength(l int64) bool {
	p.Lock()
	defer p.Unlock()
	p.MaxFileLength = l
	return p.MaxFileLength == l
}

func (p *fileWitter) Write(tab string) (n int64, err error) {

	keys, ok := Cache.Members(tab)
	if !ok {
		return 0, nil
	}

	var (
		count int
		size  int64
	)

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
			log, ok := v.(LogItem)
			if !ok {
				continue
			}
			if count, err = fileExecutor.WriteAppend(
				name, strings.Join(log.Values, log.Separator)+"\n"); err != nil {
				return
			}
			n += int64(count)
			size += int64(count)
			if size < p.MaxFileLength {
				continue
			}

			if err = p.moveFile(name); err != nil {
				return
			}
		}
	}
	return
}

func (p *fileWitter) moveFile(name string) error {
	return fileExecutor.Rename(name, name+times.TimeToDashString())
}
