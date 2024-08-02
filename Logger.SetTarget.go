package logger

// SetTarget - Set the current target of a logger object.
//
//	If multiple targets are set to true, only the right most true value
//	will be configured for logging.
func (log *Logger) SetTarget(stdout, stderr, syslog bool) (err error) {
	if stdout {
		log.UseStdout()
	}
	if stderr {
		log.UseStderr()
	}
	if syslog {
		log.UseSyslog()
	}
	return err
}
