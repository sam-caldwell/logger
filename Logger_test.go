package logger

import "testing"

func TestLogger_struct(t *testing.T) {
	var log Logger
	if log.target != nil {
		t.Fatalf("expect nil target initially")
	}
	if log.level != Critical {
		t.Fatalf("expect 0 value log level")
	}
	if log.format != "" {
		t.Fatalf("expect empty format")
	}
	log.lock.Lock()
	log.lock.Unlock()
}
