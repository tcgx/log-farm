package logfarm

const (
	// min file length is 1024 bytes
	minFileLength = 1024
)

type LoggerWritter interface {
	SetMaxFileLength(l int64) bool
	Write(tab string) (int, error)
}
