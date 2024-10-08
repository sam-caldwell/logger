package logger

import "fmt"

// Debug - write a message string to the logger.
func (log *Logger) Debug(msg string) {
	log.write(Debug, &msg)
}

// Debugf - write a formatted string to the logger.
func (log *Logger) Debugf(formatString string, msg ...any) {
	log.Debug(fmt.Sprintf(formatString, msg...))
}
