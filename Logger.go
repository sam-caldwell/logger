package logger

import (
	"sync"
)

// Logger - A simple Logger facility
type Logger struct {
	// target - The device to which log messages will be written
	target Target
	// level - log level definition consistent with RFC5424
	level Levels
	// format - the format string used to ensure uniformity among log lines
	format string
	// lock - a lock used in some cases to ensure asynchronous logging activity doesn't overlap
	lock sync.Mutex
}
