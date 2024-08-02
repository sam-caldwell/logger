package logger

// SetTarget - Set the current target of a logger object
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
