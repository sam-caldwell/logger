package logger

import "fmt"

// SetVerbosity - Set the current verbosity of a logger object
func (log *Logger) SetVerbosity(verbosity Verbosity) error {
	switch verbosity {
	case Normal, Quiet:
		log.verbosity = verbosity
	default:
		return fmt.Errorf("invalid log verbosity")
	}
	return nil
}
