package logger

import "fmt"

// Critical - write a message string to the logger.
func (log *Logger) Critical(msg string) {
	log.write(Critical, &msg)
}

// Criticalf - write a formatted string to the logger.
func (log *Logger) Criticalf(formatString, msg string) {
	log.Critical(fmt.Sprintf(formatString, msg))
}
