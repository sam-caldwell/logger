package logger

// Setup - Set up the logger object configuration.
func (log *Logger) Setup(useStdout, useStderr, useSyslog bool, level Levels) error {
	if err := log.SetTarget(useStdout, useStderr, useSyslog); err != nil {
		return err
	}
	if err := log.SetLevel(level); err != nil {
		return err
	}
	log.format = defaultLogFormat
	return nil
}
