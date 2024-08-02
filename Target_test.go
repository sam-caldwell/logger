package logger

import (
	"log/syslog"
	"os"
	"testing"
)

func TestTarget_type(t *testing.T) {
	var o Target
	o = os.Stdout
	o = os.Stderr
	o, _ = syslog.New(syslog.LOG_DEBUG, "test")
	if o == nil {
		t.Fatal("o is nil...not good")
	}
}
