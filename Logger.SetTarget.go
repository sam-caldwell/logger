package logger

import "fmt"

// SetTarget - Set the current target of a logger object
func (log *Logger) SetTarget(target Target) (err error) {
	if target == nil {
		err = fmt.Errorf("nil log target")
	} else {
		log.target = target
	}
	return err
}
