package logger

import "testing"

func TestLevels_Type(t *testing.T) {
	/*
	 *	NOTE: The order of these levels (e.g. their numeric value is extremely important).
	 */
	data := []Levels{
		Critical, Alert, Error, Warning, Notice, Info, Debug,
	}

	if Info != Informational {
		t.Fatal("info and informational should be the same value")
	}

	for i, level := range data {
		if int(level) != i {
			t.Fatalf("mismatch at %d", i)
		}
	}
}
