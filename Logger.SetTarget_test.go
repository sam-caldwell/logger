package logger

import (
	"os"
	"testing"
)

func TestLogger_SetTarget(t *testing.T) {
	t.Run("all false.  expect stdout", func(t *testing.T) {
		var log Logger
		if err := log.SetTarget(false, false, false); err != nil {
			t.Fatal(err)
		}
		if log.target != os.Stdout {
			t.Fatal("mismatched target")
		}
	})
	t.Run("false,false,true.  expect syslog", func(t *testing.T) {
		var log Logger
		if err := log.SetTarget(false, false, true); err != nil {
			t.Fatal(err)
		}
		//if log.target != goLog.Logger {
		//	t.Fatal("mismatched target")
		//}
	})
	t.Run("false, true, true.  expect stderr", func(t *testing.T) {
		var log Logger
		if err := log.SetTarget(false, true, true); err != nil {
			t.Fatal(err)
		}
		if log.target != os.Stderr {
			t.Fatal("mismatched target")
		}
	})
	t.Run("true,true,true.  expect stdout", func(t *testing.T) {
		var log Logger
		if err := log.SetTarget(true, true, true); err != nil {
			t.Fatal(err)
		}
		if log.target != os.Stdout {
			t.Fatal("mismatched target")
		}
	})
	t.Run("true,true,false.  expect stdout", func(t *testing.T) {
		var log Logger
		if err := log.SetTarget(true, true, false); err != nil {
			t.Fatal(err)
		}
		if log.target != os.Stdout {
			t.Fatal("mismatched target")
		}
	})
	t.Run("true,false,true.  expect stdout", func(t *testing.T) {
		var log Logger
		if err := log.SetTarget(true, false, true); err != nil {
			t.Fatal(err)
		}
		if log.target != os.Stdout {
			t.Fatal("mismatched target")
		}
	})
	t.Run("true,false,false.  expect stdout", func(t *testing.T) {
		var log Logger
		if err := log.SetTarget(true, false, false); err != nil {
			t.Fatal(err)
		}
		if log.target != os.Stdout {
			t.Fatal("mismatched target")
		}
	})
	t.Run("false,true,false.  expect stderr", func(t *testing.T) {
		var log Logger
		if err := log.SetTarget(false, true, false); err != nil {
			t.Fatal(err)
		}
		if log.target != os.Stderr {
			t.Fatalf("mismatched target\n"+
				"Got:  %v\n"+
				"Need: %v", log.target, os.Stderr)
		}
	})
}
