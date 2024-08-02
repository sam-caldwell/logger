package logger

import "testing"

func TestFromString(t *testing.T) {
	tests := []struct {
		input    string
		expected Levels
		err      bool
	}{
		{"critical", Critical, false},
		{"alert", Alert, false},
		{"error", Error, false},
		{"warn", Warning, false},
		{"warning", Warning, false},
		{"notice", Notice, false},
		{"info", Info, false},
		{"informational", Info, false},
		{"debug", Debug, false},
		{"  CRITICAL  ", Critical, false},
		{"invalid", 0, true},
		{"critical,debug", 0, true},
	}

	for _, tt := range tests {
		var level Levels
		err := level.FromString(tt.input)
		if (err != nil) != tt.err {
			t.Errorf("FromString(%q) error = %v, wantErr %v", tt.input, err, tt.err)
		}
		if err == nil && level != tt.expected {
			t.Errorf("FromString(%q) = %v, want %v", tt.input, level, tt.expected)
		}
	}
}
