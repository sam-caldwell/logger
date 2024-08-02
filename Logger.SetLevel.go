package logger

import "fmt"

// SetLevel - Set the current level of a logger object
func (log *Logger) SetLevel(level Levels) error {
	switch level {
	case Critical, Alert, Error, Warning, Notice, Informational, Debug:
		log.level = level
	default:
		return fmt.Errorf("invalid log level")
	}
	return nil
}
