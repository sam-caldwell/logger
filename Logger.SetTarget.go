package logger

// SetTarget - Set the current target of a logger object.
//
// This method will set the first (left-most) true target type.  Thus, if
// all three are true, stdout will take precedence.  But also if none are
// true, stdout will be the default.
func (log *Logger) SetTarget(stdout, stderr, syslog bool) (err error) {
	if stdout {
		log.UseStdout()
		return
	}
	if stderr {
		log.UseStderr()
		return
	}
	if syslog {
		log.UseSyslog()
		return
	}
	log.UseStdout()
	return err
}
