package logger

import "fmt"

// Info - write a message string to the logger.
func (log *Logger) Info(msg string) {
	log.write(Info, &msg)
}

// Infof - write a formatted string to the logger.
func (log *Logger) Infof(formatString, msg string) {
	log.Info(fmt.Sprintf(formatString, msg))
}