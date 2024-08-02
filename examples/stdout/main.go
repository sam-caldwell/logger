package main

import "github.com/sam-caldwell/logger"

func main() {
	var log logger.Logger
	log.UseStdout()
	_ = log.SetLevel(logger.Debug)
	log.Debug("test debug")
}
