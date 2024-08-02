package logger

import (
	goLog "log"
)

// write - Writes a log message to the log target if the log messageLevel satisfies the current Log level.
//
//	If messageLevel is less (higher priority) than the logging level, print the log
func (log *Logger) write(messageLevel Levels, msg *string) {
	if messageLevel <= log.level {
		if log.format == "" {
			log.format = defaultLogFormat
		}
		goLog.Printf(log.format+"\n", messageLevel.ToString(), *msg)
	}
}
