package logger

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestLogger_Debug(t *testing.T) {
	var log Logger
	verifyData := func(out string, expected string, expectEmpty bool) error {
		if !expectEmpty {
			parts := strings.Split(out, " ")
			rhs := strings.Join(parts[2:], " ")
			if rhs != expected {
				return fmt.Errorf("expected result not found\n"+
					"  Actual: '%s'\n"+
					"Expected: '%s'", rhs, expected)
			}
		}
		return nil
	}
	t.Run("Debug in Debug Level (expect log)", func(t *testing.T) {
		const expected = "{\"p\":\"   debug\",m:\"test\"}\n"
		var buffer bytes.Buffer
		log.UseBuffer(&buffer)
		if err := log.SetLevel(Debug); err != nil {
			t.Fatal(err)
		}
		log.Debug("test")
		if err := verifyData(buffer.String(), expected, false); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("Debug in Info level (expect log)", func(t *testing.T) {
		const expected = "{\"p\":\"   debug\",m:\"test\"}\n"
		var buffer bytes.Buffer
		log.UseBuffer(&buffer)
		if err := log.SetLevel(Info); err != nil {
			t.Fatal(err)
		}
		log.Debug("test")
		if err := verifyData(buffer.String(), expected, true); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("Debug in critical level (no log expected)", func(t *testing.T) {
		const expected = "{\"p\":\"   debug\",m:\"test\"}\n"
		var buffer bytes.Buffer
		log.UseBuffer(&buffer)
		if err := log.SetLevel(Critical); err != nil {
			t.Fatal(err)
		}
		log.Debug("test")
		if err := verifyData(buffer.String(), expected, true); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("formatted string: Debug in Debug level (no log expected)", func(t *testing.T) {
		const expected = "{\"p\":\"    debug\",m:\"format(test1)\"}\n"
		var buffer bytes.Buffer
		log.UseBuffer(&buffer)
		if err := log.SetLevel(Debug); err != nil {
			t.Fatal(err)
		}
		log.Debugf("format(%s%d)", "test", 1)
		if err := verifyData(buffer.String(), expected, true); err != nil {
			t.Fatal(err)
		}
	})
}
