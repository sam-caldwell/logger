package main

import (
	"flag"
	"github.com/sam-caldwell/logger"
	"os"
)

func main() {
	var log logger.Logger
	log.UseStdout()
	if err := log.SetLevel(logger.Critical); err != nil {
		os.Exit(255)
	}
	exitCode := flag.Int("exit", 0, "exit code")
	message := flag.String("msg", "", "log message")
	flag.Parse()
	log.Fatal(*message, *exitCode)
}
