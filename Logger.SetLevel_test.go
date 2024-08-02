package logger

import (
	"fmt"
	"testing"
)

func TestLogger_SetLevel(t *testing.T) {
	t.Run("test valid points", func(t *testing.T) {
		var log Logger

		data := []Levels{Critical, Alert, Error, Warning, Notice, Informational, Debug}

		for _, level := range data {
			t.Run(fmt.Sprintf("Test %v", level), func(t *testing.T) {
				if err := log.SetLevel(level); err != nil {
					t.Fatal(err)
				}
				if log.level != level {
					t.Fatalf("mismatch on %v", level)
				}
			})
		}
	})
	t.Run("test invalid point", func(t *testing.T) {
		var log Logger
		if err := log.SetLevel(7); err == nil {
			t.Fatal("invalid level expects error")
		} else {
			if err.Error() != "invalid log level" {
				t.Fatal("expect 'invalid log level' error message")
			}
		}
	})
}
