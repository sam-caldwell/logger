package logger

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestLogger_Critical(t *testing.T) {
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
	t.Run("Critical in Error Level (expect log)", func(t *testing.T) {
		const expected = "{\"p\":\"critical\",m:\"test\"}\n"
		var buffer bytes.Buffer
		log.UseBuffer(&buffer)
		if err := log.SetLevel(Error); err != nil {
			t.Fatal(err)
		}
		log.Critical("test")
		if err := verifyData(buffer.String(), expected, true); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("Critical in Alert level (expect log)", func(t *testing.T) {
		const expected = "{\"p\":\"critical\",m:\"test\"}\n"
		var buffer bytes.Buffer
		log.UseBuffer(&buffer)
		if err := log.SetLevel(Alert); err != nil {
			t.Fatal(err)
		}
		log.Critical("test")
		if err := verifyData(buffer.String(), expected, true); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("Critical in critical level (no log expected)", func(t *testing.T) {
		const expected = "{\"p\":\"critical\",m:\"test\"}\n"
		var buffer bytes.Buffer
		log.UseBuffer(&buffer)
		if err := log.SetLevel(Critical); err != nil {
			t.Fatal(err)
		}
		log.Critical("test")
		if err := verifyData(buffer.String(), expected, false); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("formatted string: Critical in Critical level (no log expected)", func(t *testing.T) {
		const expected = "{\"p\":\"    critical\",m:\"format(test1)\"}\n"
		var buffer bytes.Buffer
		log.UseBuffer(&buffer)
		if err := log.SetLevel(Info); err != nil {
			t.Fatal(err)
		}
		log.Criticalf("format(%s%d)", "test", 1)
		if err := verifyData(buffer.String(), expected, true); err != nil {
			t.Fatal(err)
		}
	})
}
