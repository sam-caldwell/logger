package logger

import (
	"errors"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestLogger_Setup(t *testing.T) {
	type TestData struct {
		Stdout bool
		Stderr bool
		Syslog bool
		Level  Levels
		Error  error
		Writer io.Writer
	}
	testData := []TestData{
		{false, false, false, Debug, nil, os.Stdout},
		{true, false, false, Debug, nil, os.Stdout},
		{false, true, false, Debug, nil, os.Stderr},
		{false, false, false, Informational, nil, os.Stdout},
		{true, false, false, Informational, nil, os.Stdout},
		{false, true, false, Informational, nil, os.Stderr},
		{false, false, false, Info, nil, os.Stdout},
		{true, false, false, Info, nil, os.Stdout},
		{false, true, false, Info, nil, os.Stderr},
		{false, false, false, Critical, nil, os.Stdout},
		{true, false, false, Critical, nil, os.Stdout},
		{false, true, false, Critical, nil, os.Stderr},
		{false, false, false, Alert, nil, os.Stdout},
		{true, false, false, Alert, nil, os.Stdout},
		{false, true, false, Alert, nil, os.Stderr},
	}

	t.Run("test setup with test data", func(t *testing.T) {
		for row, data := range testData {
			t.Run(fmt.Sprintf("row %d", row), func(t *testing.T) {
				var log Logger
				if err := log.Setup(data.Stdout, data.Stderr, data.Syslog, data.Level); !errors.Is(err, data.Error) {
					t.Fatal(err)
				}
			})
		}
	})
}
