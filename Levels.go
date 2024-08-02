package logger

// Levels - A numeric indicator of the logging level
//
//	The assigned values follows https://datatracker.ietf.org/doc/html/rfc5424
type Levels uint8

const (
	// Critical - system (application) is unusable
	Critical Levels = 0
	// Alert - action must be taken immediately
	Alert Levels = 1
	// Error - error conditions exist
	Error Levels = 2
	// Warning - warning conditions exist
	Warning Levels = 3
	// Notice - normal but significant conditions
	Notice Levels = 4
	// Info - informational messages
	Info Levels = 5
	// Informational - longer form of Info
	Informational Levels = 5
	// Debug - Debugging information
	Debug Levels = 6
)
