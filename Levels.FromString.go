package logger

import (
	"fmt"
	"strings"
)

// FromString - Given a string input, convert and store the matching log level or return an error
func (level *Levels) FromString(levelString string) error {
	switch strings.TrimSpace(strings.ToLower(levelString)) {
	case "critical":
		*level = Critical
	case "alert":
		*level = Alert
	case "error":
		*level = Error
	case "warn", "warning":
		*level = Warning
	case "notice":
		*level = Notice
	case "info", "informational":
		*level = Info
	case "debug":
		*level = Debug
	default:
		return fmt.Errorf("invalid log level string")
	}
	return nil
}
