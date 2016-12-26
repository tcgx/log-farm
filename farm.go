package logfarm

type LogFarm interface {
	SetLevel(Level)
	SetLogFilename(name string)

	WriteLog(context string)
	WriteLogIntoFilename(name, context string)

	Debug(...interface{})
	Info(...interface{})
	Warn(...interface{})
	Error(...interface{})
	Panic(...interface{})
}
