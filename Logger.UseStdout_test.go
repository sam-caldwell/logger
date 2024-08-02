package logger

import (
	"bytes"
	"os/exec"
	"strings"
	"testing"
)

func TestLogger_UseStdout(t *testing.T) {
	var log Logger
	t.Run("configure stdout", func(t *testing.T) {
		log.UseStdout()
	})

	t.Run("call examples/stdout/main.go as test program", func(t *testing.T) {
		// Run the example program and capture stdout
		cmd := exec.Command("go", "run", "examples/stdout/main.go")
		var stdout bytes.Buffer
		cmd.Stdout = &stdout

		// Execute the command
		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}

		// Capture the output from stderr
		out := stdout.String()

		parts := strings.Split(out, " ")
		rhs := strings.Join(parts[2:], " ")
		const expected = "{\"p\":\"   debug\",m:\"test debug\"}\n"
		if rhs != expected {
			t.Fatalf("expected result not found\n"+
				"  Actual: '%s'\n"+
				"Expected: '%s'", rhs, expected)
		}
	})
}
