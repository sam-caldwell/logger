package logger

import (
	goLog "log"
	"time"
)

// write - Writes a log message to the log target if the log messageLevel satisfies the current Log level.
func (log *Logger) write(messageLevel Levels, msg *string) {

	if log.verbosity == Quiet {
		return //Print nothing
	}

	// if messageLevel is less (higher priority) than the logging level, print the log
	if messageLevel <= log.level {
		//Print the actual message.
		goLog.Printf(log.format+"\n", time.Now(), messageLevel.ToString(), *msg)
	}

}
