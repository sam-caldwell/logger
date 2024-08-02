package logger

import "os"

// UseStderr - Configure logger for stderr
func (log *Logger) UseStderr() {
	log.target = os.Stderr
}
