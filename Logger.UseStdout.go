package logger

import (
	goLog "log"
	"os"
)

// UseStdout - configure logger for stdout
func (log *Logger) UseStdout() {
	log.target = os.Stdout
	goLog.SetOutput(log.target)

}
