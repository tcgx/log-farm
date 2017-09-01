# log-farm
a tool to write logs to file with format data by separator
---

## Introduce

```go
// LogFarm functions to wite logs
type LogFarm interface {
	// Write log into cache
	WriteLog(data []string) bool
	// stop write data into file
	Stop()
}

logfarm.New(filename, config.Options)
```

* filename: the log path & name, see more in example
* filesuffix: log file' suffix
* chanbuffer: length of the log chan buffer
* filemaxlength: the max length of log file, default: 0 is unlimit
* movefiletype: move file by per-minite(1) or hourly(2) or daily(3), 0 is doing nothing

## [example](example/main.go)
