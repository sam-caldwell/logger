package logger

import (
	"os/exec"
	"runtime"
	"strings"
	"testing"
)

func TestLogger_UseSyslog(t *testing.T) {
	var log Logger

	filterLines := func(lines []string) (result []string) {
		for _, line := range lines {
			if strings.Contains(line, "myapp") {
				result = append(result, line)
			}
		}
		return result
	}

	getLogs := func() (lines []string, count int) {
		if runtime.GOOS != "linux" {
			t.Fatal("test only supports linux.  If you see this, fix the test for your operating system")
		}
		out, err := exec.Command("tail", "-n100", "/var/log/syslog").Output()
		if err != nil {
			t.Fatal(err)
		}
		lines = strings.Split(string(out), "\n")
		lines = filterLines(lines)
		return lines, len(lines)
	}

	t.Run("configure stdout", func(t *testing.T) {
		log.UseSyslog()
	})

	t.Run("call examples/syslog/main.go as test program", func(t *testing.T) {
		_, startCount := getLogs()
		// Run the example program and capture stdout
		out, err := exec.Command("go", "run", "examples/syslog/main.go").Output()
		if err != nil {
			t.Fatal(err)
		}
		if string(out) != "" {
			t.Fatalf("expected no stdout output\n"+
				"got: '%v'", out)
		}
		lines, endCount := getLogs()
		if endCount-startCount != 1 {
			t.Fatalf("count mismatch.\n"+
				"  end: %d\n"+
				"start: %d\n"+
				"lines: %s", endCount, startCount, lines)
		}
		lastLine := lines[len(lines)-1]
		parts := strings.Split(lastLine, " ")
		rhs := strings.Join(parts[8:], " ")
		const expected = "{\"p\":\"   debug\",m:\"test debug\"}"
		if rhs != expected {
			t.Fatalf("expected result not found\n"+
				"  Actual: '%s'\n"+
				"Expected: '%s'", rhs, expected)
		}
	})
}
