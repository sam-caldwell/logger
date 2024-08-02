package logger

import "testing"

func TestGetListOfLevels(t *testing.T) {
	levels := GetListOfLevels()
	if levels[0] != "critical" {
		t.Fatalf("failed at critical. Got: %v", levels[0])
	}
	if levels[1] != "alert" {
		t.Fatalf("failed at alert. Got: %v", levels[1])
	}
	if levels[2] != "error" {
		t.Fatalf("failed at error. Got: %v", levels[2])
	}
	if levels[3] != "warn" {
		t.Fatalf("failed at warning. Got: %v", levels[3])
	}
	if levels[4] != "notice" {
		t.Fatalf("failed at notice. Got: %v", levels[4])
	}
	if levels[5] != "info" {
		t.Fatalf("failed at info. Got: %v", levels[5])
	}
	if levels[6] != "debug" {
		t.Fatalf("failed at debug. Got: %v", levels[6])
	}
}
