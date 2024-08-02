package logger

import (
	goLog "log"
	"log/syslog"
)

// UseSyslog - Configure logger for syslog
func (log *Logger) UseSyslog() {
	var err error
	p := log.level.ToSysLogPriority()
	if log.target, err = syslog.New(p, "myapp"); err != nil {
		goLog.Fatal("failed to connect to syslog:", err)
	}
	goLog.SetOutput(log.target)
}
