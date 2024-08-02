package logger

import (
	"bytes"
	"os/exec"
	"strings"
	"testing"
)

func TestLogger_UseStderr(t *testing.T) {
	var log Logger
	t.Run("configure stderr", func(t *testing.T) {
		log.UseStderr()
	})

	t.Run("call examples/stderr/main.go as test program", func(t *testing.T) {
		// Run the example program and capture stderr
		cmd := exec.Command("go", "run", "examples/stderr/main.go")
		var stderr bytes.Buffer
		cmd.Stderr = &stderr

		// Execute the command
		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}

		// Capture the output from stderr
		out := stderr.String()

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
