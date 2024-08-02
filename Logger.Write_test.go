package logger

import (
	"bytes"
	"strings"
	"testing"
)

func TestLogger_Write(t *testing.T) {
	t.Run("test write with debug|info", func(t *testing.T) {
		var log Logger
		var buffer bytes.Buffer
		msg := "test"
		log.UseBuffer(&buffer)
		_ = log.SetLevel(Debug)
		log.write(Info, &msg)
		out := buffer.String()

		parts := strings.Split(out, " ")
		rhs := strings.Join(parts[2:], " ")
		const expected = "{\"p\":\"    info\",m:\"test\"}\n"
		if rhs != expected {
			t.Fatalf("expected result not found\n"+
				"  Actual: '%s'\n"+
				"Expected: '%s'", rhs, expected)
		}
	})
	t.Run("test write with debug|debug", func(t *testing.T) {
		var log Logger
		var buffer bytes.Buffer
		msg := "test"
		log.UseBuffer(&buffer)
		_ = log.SetLevel(Debug)
		log.write(Debug, &msg)
		out := buffer.String()

		parts := strings.Split(out, " ")
		rhs := strings.Join(parts[2:], " ")
		const expected = "{\"p\":\"   debug\",m:\"test\"}\n"
		if rhs != expected {
			t.Fatalf("expected result not found\n"+
				"  Actual: '%s'\n"+
				"Expected: '%s'", rhs, expected)
		}
	})

	t.Run("test write with debug|critical", func(t *testing.T) {
		var log Logger
		var buffer bytes.Buffer
		msg := "test"
		log.UseBuffer(&buffer)
		_ = log.SetLevel(Debug)
		log.write(Critical, &msg)
		out := buffer.String()

		parts := strings.Split(out, " ")
		rhs := strings.Join(parts[2:], " ")
		const expected = "{\"p\":\"critical\",m:\"test\"}\n"
		if rhs != expected {
			t.Fatalf("expected result not found\n"+
				"  Actual: '%s'\n"+
				"Expected: '%s'", rhs, expected)
		}
	})

	t.Run("test write with alert|debug", func(t *testing.T) {
		var log Logger
		var buffer bytes.Buffer
		msg := "test"
		log.UseBuffer(&buffer)
		_ = log.SetLevel(Alert)
		log.write(Debug, &msg)
		out := buffer.String()
		if out != "" {
			t.Fatal("expected no log.  debug is not high enough priority to write during alert level")
		}
	})

	t.Run("test write with alert|error", func(t *testing.T) {
		var log Logger
		var buffer bytes.Buffer
		msg := "test"
		log.UseBuffer(&buffer)
		_ = log.SetLevel(Alert)
		log.write(Error, &msg)
		out := buffer.String()
		if out != "" {
			t.Fatal("expected no log.  error is not high enough priority to write during alert level")
		}
	})

	t.Run("test write with debug|critical", func(t *testing.T) {
		var log Logger
		var buffer bytes.Buffer
		msg := "test"
		log.UseBuffer(&buffer)
		_ = log.SetLevel(Alert)
		log.write(Critical, &msg)
		out := buffer.String()

		parts := strings.Split(out, " ")
		rhs := strings.Join(parts[2:], " ")
		const expected = "{\"p\":\"critical\",m:\"test\"}\n"
		if rhs != expected {
			t.Fatalf("expected result not found\n"+
				"  Actual: '%s'\n"+
				"Expected: '%s'", rhs, expected)
		}
	})
}
