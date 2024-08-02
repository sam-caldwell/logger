package logger

import "fmt"

// Notice - write a message string to the logger.
func (log *Logger) Notice(msg string) {
	log.write(Notice, &msg)
}

// Noticef - write a formatted string to the logger.
func (log *Logger) Noticef(formatString, msg string) {
	log.Notice(fmt.Sprintf(formatString, msg))
}