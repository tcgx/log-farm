package main

import (
	"fmt"
	"time"

	"github.com/go-trellis/config"
	"github.com/go-trellis/log-farm"
)

var (
	loggerTest1 = logfarm.New("log/test1.log", nil)
	loggerTest2 = logfarm.New("log/test2",
		config.Options{
			"chanbuffer":  1000,
			"filesuffix":   "log",
			"movefiletype": 1})
	loggerTest3 = logfarm.New("log/test3",
		config.Options{
			"filemaxlength": 10240,
			"chanbuffer":    1000,
			"filesuffix":    "log",
			"movefiletype":  3,
			"separator":     "::"})
)

func main() {

	for i := 0; i < 10000; i++ {
		if (i+1)%1000 == 0 {
			fmt.Println(i+1, "over")
		}
		loggerTest1.WriteLog([]string{time.Now().Format("20060102150405"), "1", "2", "3", "4", "5"})
	}

	for i := 0; i < 100000; i++ {
		if (i+1)%10000 == 0 {
			fmt.Println(i+1, "over")
		}
		loggerTest2.WriteLog([]string{time.Now().Format("20060102150405"), "1", "2", "3", "4", "5", "6"})
	}

	for i := 0; i < 30000; i++ {
		if (i+1)%10000 == 0 {
			fmt.Println(i+1, "over")
		}
		loggerTest3.WriteLog([]string{time.Now().Format("20060102150405"), "1", "2", "3", "4", "5", "6", "7"})
	}

	loggerTest1.Stop()
	loggerTest2.Stop()
	loggerTest3.Stop()
	time.Sleep(time.Second * 1)
	fmt.Println("...shell: wc *log*")

}
