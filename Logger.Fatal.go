package logger

import (
	"fmt"
	"os"
)

// Fatal - write a message string to the logger. Then terminate.
func (log *Logger) Fatal(msg string, exitCode int) {
	log.write(Critical, &msg)
	os.Exit(exitCode)
}

// Fatalf - write a formatted string to the logger. Then terminate.
func (log *Logger) Fatalf(formatString, msg string, exitCode int) {
	log.Fatal(fmt.Sprintf(formatString, msg), exitCode)
}
