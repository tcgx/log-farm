package logfarm

const (
	// min length is 102400 bytes
	minLength int64 = 102400
)

// LoggerWritter
type LoggerWritter interface {
	//
	SetMaxLength(l int64) bool
	Write(tab string) (int64, error)
	ResetTab(tab string) bool
}
