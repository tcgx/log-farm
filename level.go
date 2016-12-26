package logfarm

import "github.com/go-rut/errors"

// Level defines
type Level uint8

const (
	LevelUnknown Level = iota
	LevelPanic
	LevelFatal
	LevelError
	LevelWarn
	LevelInfo
	LevelDebug
)

const (
	levelStringUnknown = "unknown"
	levelStringPanic   = "panic"
	levelStringFatal   = "fatal"
	levelStringError   = "error"
	levelStringWarn    = "warn"
	levelStringInfo    = "info"
	levelStringDebug   = "debug"
)

func (p Level) String() string {
	switch p {
	case LevelPanic:
		return levelStringPanic
	case LevelFatal:
		return levelStringFatal
	case LevelError:
		return levelStringError
	case LevelWarn:
		return levelStringWarn
	case LevelInfo:
		return levelStringInfo
	case LevelDebug:
		return levelStringDebug
	default:
		return levelStringUnknown
	}
}

func StringLevel(lvl string) (Level, error) {

	switch lvl {
	case levelStringPanic:
		return LevelPanic, nil
	case levelStringFatal:
		return LevelFatal, nil
	case levelStringError:
		return LevelError, nil
	case levelStringWarn:
		return LevelWarn, nil
	case levelStringInfo:
		return LevelInfo, nil
	case levelStringDebug:
		return LevelDebug, nil
	}

	return LevelUnknown, ErrUnknownStringLevel.New(errors.Params{"level": lvl})
}
