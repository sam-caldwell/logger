package logger

import (
	goLog "log"
	"time"
)

// write - Writes a log message to the log target if the log messageLevel satisfies the current Log level.
//
//	If messageLevel is less (higher priority) than the logging level, print the log
func (log *Logger) write(messageLevel Levels, msg *string) {
	if messageLevel <= log.level {
		goLog.Printf(log.format+"\n", time.Now(), messageLevel.ToString(), *msg)
	}
}
