package logger

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestLogger_Critical(t *testing.T) {
	var log Logger

	verifyData := func(out string, expectEmpty bool) error {
		if expectEmpty {

		} else {

			parts := strings.Split(out, " ")
			rhs := strings.Join(parts[2:], " ")
			const expected = "{\"p\":\"critical\",m:\"test\"}\n"
			if rhs != expected {
				return fmt.Errorf("expected result not found\n"+
					"  Actual: '%s'\n"+
					"Expected: '%s'", rhs, expected)
			}
		}
		return nil
	}

	t.Run("Alert in Error Level (expect log)", func(t *testing.T) {
		var buffer bytes.Buffer
		log.UseBuffer(&buffer)
		if err := log.SetLevel(Error); err != nil {
			t.Fatal(err)
		}
		log.Critical("test")
		if err := verifyData(buffer.String(), true); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("Alert in Alert level (expect log)", func(t *testing.T) {
		var buffer bytes.Buffer
		log.UseBuffer(&buffer)
		if err := log.SetLevel(Alert); err != nil {
			t.Fatal(err)
		}
		log.Critical("test")
		if err := verifyData(buffer.String(), true); err != nil {
			t.Fatal(err)
		}

	})
	t.Run("Alert in critical level (no log expected)", func(t *testing.T) {
		var buffer bytes.Buffer
		log.UseBuffer(&buffer)
		if err := log.SetLevel(Critical); err != nil {
			t.Fatal(err)
		}
		log.Critical("test")
		if err := verifyData(buffer.String(), false); err != nil {
			t.Fatal(err)
		}

	})
}
