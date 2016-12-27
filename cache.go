package logfarm

import (
	"github.com/go-trellis/cache"
)

var (
	Cache = cache.New()
)

const (
	VerA = namespace + "::verA"
	VerB = namespace + "::verB"
	VerC = namespace + "::verC"
)

func init() {
	if e := Cache.New(VerA, cache.TableOptions{
		cache.TableOptionValueMode: cache.ValueModeDuplicateBag}); e != nil {
		panic(e)
	}
	if e := Cache.New(VerB, cache.TableOptions{
		cache.TableOptionValueMode: cache.ValueModeDuplicateBag}); e != nil {
		panic(e)
	}
	if e := Cache.New(VerC, cache.TableOptions{
		cache.TableOptionValueMode: cache.ValueModeDuplicateBag}); e != nil {
		panic(e)
	}
}
