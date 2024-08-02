logger
======

## Description

A simple logger facility for golang.  This logger supports logging to stdout, stderr and the local system's logging
facility (e.g. syslog).

## Usage
```go
package main

import "github.com/sam-caldwell/logger"

func main(){
    var log Logger
    const (
        logTarget logger.LogTarget = logger.Syslog
		verbosity logger.Verbosity = logger.Quiet //as opposed to logger.Debug
		logLevel  logger.Levels = logger.Info
    )
    log.Setup(logTarget, verbosity, logLevel)
	log.Debug("Debug message")
	log.Informational("Informational message") 
	log.Notice("Notice message")
	log.Warning("Warning message")
	log.Error("Error message")
	log.Alert("Alert error message")
	log.Critical("Critical error message.")
}

```
