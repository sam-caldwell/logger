package logger

import (
	goLog "log"
	"log/syslog"
)

// UseSyslog - Configure logger for syslog
func (log *Logger) UseSyslog() {
	var err error
	if log.target, err = syslog.New(syslog.LOG_NOTICE|syslog.LOG_USER, "myapp"); err != nil {
		goLog.Fatal("failed to connect to syslog:", err)
	}
}
