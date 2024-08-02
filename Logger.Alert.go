package logger

import "fmt"

// Alert - write a message string to the logger.
func (log *Logger) Alert(msg string) {
	log.write(Alert, &msg)
}

// Alertf - write a formatted string to the logger.
func (log *Logger) Alertf(formatString, msg string) {
	log.Alert(fmt.Sprintf(formatString, msg))
}
