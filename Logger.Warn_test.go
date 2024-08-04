package logger

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestLogger_Warn(t *testing.T) {
	var log Logger

	verifyData := func(out string, expected string, expectEmpty bool) error {
		if expectEmpty {

		} else {
			parts := strings.Split(out, " ")
			rhs := strings.Join(parts[2:], " ")
			const expected = "{\"p\":\"    warn\",m:\"test\"}\n"
			if rhs != expected {
				return fmt.Errorf("expected result not found\n"+
					"  Actual: '%s'\n"+
					"Expected: '%s'", rhs, expected)
			}
		}
		return nil
	}
	t.Run("Warn in Warn level (expect log)", func(t *testing.T) {
		const expected = "{\"p\":\"    warn\",m:\"test\"}\n"
		var buffer bytes.Buffer
		log.UseBuffer(&buffer)
		if err := log.SetLevel(Warning); err != nil {
			t.Fatal(err)
		}
		log.Warn("test")
		if err := verifyData(buffer.String(), expected, false); err != nil {
			t.Fatal(err)
		}

	})
	t.Run("Warn in Error Level (expect log)", func(t *testing.T) {
		const expected = "{\"p\":\"    warn\",m:\"test\"}\n"
		var buffer bytes.Buffer
		log.UseBuffer(&buffer)
		if err := log.SetLevel(Error); err != nil {
			t.Fatal(err)
		}
		log.Warn("test")
		if err := verifyData(buffer.String(), expected, true); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("Warn in critical level (no log expected)", func(t *testing.T) {
		const expected = "{\"p\":\"    warn\",m:\"test\"}\n"
		var buffer bytes.Buffer
		log.UseBuffer(&buffer)
		if err := log.SetLevel(Critical); err != nil {
			t.Fatal(err)
		}
		log.Warn("test")
		if err := verifyData(buffer.String(), expected, true); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("formatted string: Warn in Warn level (no log expected)", func(t *testing.T) {
		const expected = "{\"p\":\"    warn\",m:\"format(test1)\"}\n"
		var buffer bytes.Buffer
		log.UseBuffer(&buffer)
		if err := log.SetLevel(Warning); err != nil {
			t.Fatal(err)
		}
		log.Warnf("format(%s%d)", "test", 1)
		if err := verifyData(buffer.String(), expected, true); err != nil {
			t.Fatal(err)
		}
	})
}
