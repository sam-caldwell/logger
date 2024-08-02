package logger

// Verbosity - a numeric value indicating how verbose the logging should be (quiet or normal)
type Verbosity uint8

const (
	Normal Verbosity = 0
	Quiet  Verbosity = 1
)
