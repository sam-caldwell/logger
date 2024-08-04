package logger

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestLogger_Notice(t *testing.T) {
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
	t.Run("Notice in Notice level (expect log)", func(t *testing.T) {
		const expected = "{\"p\":\"  notice\",m:\"test\"}\n"
		var buffer bytes.Buffer
		log.UseBuffer(&buffer)
		if err := log.SetLevel(Notice); err != nil {
			t.Fatal(err)
		}
		log.Notice("test")
		if err := verifyData(buffer.String(), expected, false); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("Notice in Error Level (expect log)", func(t *testing.T) {
		const expected = "{\"p\":\"  notice\",m:\"test\"}\n"
		var buffer bytes.Buffer
		log.UseBuffer(&buffer)
		if err := log.SetLevel(Error); err != nil {
			t.Fatal(err)
		}
		log.Notice("test")
		if err := verifyData(buffer.String(), expected, true); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("formatted string: Notice in Notice level (no log expected)", func(t *testing.T) {
		const expected = "{\"p\":\"    notice\",m:\"format(test1)\"}\n"
		var buffer bytes.Buffer
		log.UseBuffer(&buffer)
		if err := log.SetLevel(Notice); err != nil {
			t.Fatal(err)
		}
		log.Noticef("format(%s%d)", "test", 1)
		if err := verifyData(buffer.String(), expected, true); err != nil {
			t.Fatal(err)
		}
	})
}
