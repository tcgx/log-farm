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
			"chanbuffer":   1000,
			"filesuffix":   "log",
			"movefiletype": 1})
	loggerTest3 = logfarm.New("log/test3",
		config.Options{
			"filemaxlength": 102400,
			"chanbuffer":    1000,
			"filesuffix":    "log",
			"movefiletype":  3,
			"separator":     "::"})
)

func main() {

	writer("log1", loggerTest1, 10000, "1", "2", "3", "4", "5")

	writer("log2", loggerTest2, 100000, "1", "2", "3", "4", "5", "6")

	writer("log3", loggerTest3, 30000, "1", "2", "3", "4", "5", "6", "7")

	time.Sleep(time.Second * 60)
	loggerTest1.Stop()
	loggerTest2.Stop()
	loggerTest3.Stop()
	fmt.Println("...shell: wc *log*")
}

func writer(test string, log logfarm.LogFarm, times int, data ...string) {

	for i := 0; i < times; i++ {
		if (i+1)%10000 == 0 {
			fmt.Println(test, i+1, "over")
		}
		log.WriteLog(append([]string{time.Now().Format("20060102150405")}, data...))
	}
}
