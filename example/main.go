package main

import (
	"fmt"
	"time"

	"github.com/go-trellis/log-farm"

	"github.com/go-trellis/files"
)

var (
	logger = logfarm.New()
	fi     = files.New()
)

func main() {
	// logger.SetSeparator("::")
	logger.SetLoopTimerToWriteLog(time.Second * 2)

	for i := 0; i < 1000; i++ {
		logger.WriteLog("test1.log", []string{"1", "2", "3", "4", "5"})
	}

	time.Sleep(time.Second * 3)

	for i := 0; i < 10000; i++ {
		logger.WriteLog("test2.log", []string{"1", "2", "3", "4", "5"})
	}

	time.Sleep(time.Second * 10)

	fmt.Println("...shell: wc *log*")

}
