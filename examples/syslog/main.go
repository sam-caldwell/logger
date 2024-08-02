package main

import "github.com/sam-caldwell/logger"

func main() {
	var log logger.Logger
	log.UseSyslog()
	_ = log.SetLevel(logger.Debug)
	log.Debug("test debug")
}
