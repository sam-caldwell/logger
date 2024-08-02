package logger

import "os"

// UseStdout - configure logger for stdout
func (log *Logger) UseStdout() {
	log.target = os.Stdout
}
