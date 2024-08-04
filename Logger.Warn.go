package logger

import "fmt"

// Warn - write a message string to the logger.
func (log *Logger) Warn(msg string) {
	log.write(Warning, &msg)
}

// Warnf - write a formatted string to the logger.
func (log *Logger) Warnf(formatString string, msg ...any) {
	log.Warn(fmt.Sprintf(formatString, msg...))
}
