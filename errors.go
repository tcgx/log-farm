package logfarm

import "github.com/go-rut/errors"

const (
	namespace = "Trellis::LogFarm"
)

var (
	ErrUnknownStringLevel = errors.TN(namespace, 1000, "level {{.level}} is unknown")
)
