package main

import (
	"github.com/sam-caldwell/logger"
	"os"
)

func main() {
	var log logger.Logger
	log.UseStdout()
	log.UseStderr()
	log.UseSyslog()
	_ = log.SetLevel(logger.Informational)
	_ = log.SetTarget(true, true, true)
	log.Critical("A critical message")
	log.Alert("An alert")
	log.Error("An error")
	log.Warn("a warning")
	log.Notice("a notice")
	log.Info("an info message")
	log.Debug("a debug message")
	log.Fatal("a fatal message", 0)
	//If the .Fatal() call didn't fail with exit code 0, we will fail with 1
	os.Exit(1)
}
