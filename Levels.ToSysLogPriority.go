package logger

import "log/syslog"

// ToSyslogPriority - Convert our log levels to the syslog priorities.
//
// The internal levels follow the syslog RFC pretty closely.  But
// these levels may add non-RFC values in the future.  This method
// converts our log levels to an RFC5424 user level priority value.
// See https://datatracker.ietf.org/doc/html/rfc5424
func (level *Levels) ToSyslogPriority() syslog.Priority {
	switch *level {
	case Critical:
		return syslog.LOG_USER | syslog.LOG_CRIT
	case Alert:
		return syslog.LOG_USER | syslog.LOG_ALERT
	case Error:
		return syslog.LOG_USER | syslog.LOG_ERR
	case Warning:
		return syslog.LOG_USER | syslog.LOG_WARNING
	case Notice:
		return syslog.LOG_USER | syslog.LOG_NOTICE
	case Informational:
		return syslog.LOG_USER | syslog.LOG_INFO
	case Debug:
		return syslog.LOG_USER | syslog.LOG_DEBUG
	default:
		panic("invalid priority")
	}
}
