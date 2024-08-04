package logger

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestLogger_Fatal(t *testing.T) {
	const testBinary = "./example_fatal_test_program"
	t.Cleanup(func() {
		if err := os.Remove(testBinary); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("test build", func(t *testing.T) {
		out, err := exec.Command("go", "build", "-o", testBinary, "examples/fatal/main.go").Output()
		if err != nil {
			t.Fatalf("error building test binary\n"+
				"out: %v\n"+
				"err: %v", out, err)
		}
	})
	t.Run("Fatal in Critical Level (expect log/exitCode)", func(t *testing.T) {
		n := 10
		msg := "test"
		currentLevel := Critical
		exitCode := fmt.Sprintf("%d", n)
		cmd := exec.Command(testBinary, "-msg", msg, "-exit", exitCode)
		var stdout, stderr bytes.Buffer
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr
		err := cmd.Run()
		t.Run("Test exit code", func(t *testing.T) {
			// Check if the command failed (expected since it calls os.Exit)
			if exitError, ok := err.(*exec.ExitError); ok {
				if exitError.ExitCode() != n {
					t.Fatalf("expected exit code\n"+
						"want %d ('%s')\n"+
						" got %d", n, exitCode, exitError.ExitCode())
				}
			} else if err != nil {
				t.Fatalf("cmd.Run() failed with %s", err)
			}
		})

		t.Run("Test stderr", func(t *testing.T) {
			errOut := stderr.String()
			if errOut != "" {
				t.Fatalf("unexpected output on stderr: %v", stderr)
			}
		})
		t.Run("test stdout", func(t *testing.T) {
			// Capture the output
			out := stdout.String()
			t.Log(out)

			parts := strings.Split(out, " ")
			if len(parts) < 3 {
				t.Fatal("output should have at least 3 parts")
			}
			actual := strings.Join(parts[2:], " ")
			// Log the outputs for debugging purposes
			t.Logf("stdout: '%s'", actual)

			// Check if the stdout matches the expected message

			expected := fmt.Sprintf("{\"p\":\"%8s\",m:\"%s\"}\n", currentLevel.ToString(), msg)
			if actual != expected {
				t.Fatalf("stdout mismatch\n"+
					"actual:'%s'\n"+
					"expect:'%s'", actual, expected)
			}
		})
	})
}
