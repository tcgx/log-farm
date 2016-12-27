// GNU GPL v3 License

// Copyright (c) 2016 github.com:go-trellis

package logfarm

import (
	"github.com/go-trellis/cache"
)

var Cache = cache.New()

const (
	VerA = namespace + "::verA"
	VerB = namespace + "::verB"
	VerC = namespace + "::verC"
)

func init() {
	if e := initCacheTabels(VerA, VerB, VerC); e != nil {
		panic(e)
	}
}

func initCacheTabels(tabs ...string) (err error) {
	for _, v := range tabs {
		if err = Cache.New(v,
			cache.TableOptions{
				cache.TableOptionValueMode: cache.ValueModeDuplicateBag,
			}); err != nil {
			return
		}
	}
	return nil
}
