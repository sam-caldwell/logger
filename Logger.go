package logger

import (
	"sync"
)

// Logger - A simple Logger facility
type Logger struct {
	target    Target
	verbosity Verbosity
	level     Levels
	format    string
	lock      sync.Mutex
}
