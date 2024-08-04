package logger

import "fmt"

// Error - write a message string to the logger.
func (log *Logger) Error(msg string) {
	log.write(Error, &msg)
}

// Errorf - write a formatted string to the logger.
func (log *Logger) Errorf(formatString string, msg ...any) {
	log.Error(fmt.Sprintf(formatString, msg...))
}
