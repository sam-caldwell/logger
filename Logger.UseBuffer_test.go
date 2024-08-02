package logger

import (
	"bytes"
	"testing"
)

func TestLogger_UseBuffer(t *testing.T) {
	var log Logger
	t.Run("nil writer expects error", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Fatalf("expected error not found")
			} else {
				if r != "cannot use nil writer" {
					t.Fatalf("unexpected error message")
				}
			}
		}()
		log.UseBuffer(nil)
		t.Fatalf("expected panic")
	})

	t.Run("non-nil writer", func(t *testing.T) {
		var buffer bytes.Buffer
		log.UseBuffer(&buffer)
	})
}
