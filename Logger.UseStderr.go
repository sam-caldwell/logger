package logger

import (
	goLog "log"
	"os"
)

// UseStderr - Configure logger for stderr
func (log *Logger) UseStderr() {
	log.target = os.Stderr
	goLog.SetOutput(log.target)
}
