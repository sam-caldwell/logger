package logger

func (level *Levels) ToString() string {
	switch *level {
	case Critical:
		return "critical"
	case Alert:
		return "alert"
	case Error:
		return "error"
	case Warning:
		return "warn"
	case Notice:
		return "notice"
	case Informational:
		return "info"
	case Debug:
		return "debug"
	default:
		panic("invalid log level")
	}
}
