package logger

// GetListOfLevels - return a []string list of log levels
func GetListOfLevels() (result []string) {
	levels := []Levels{Critical, Alert, Error, Warning, Notice, Info, Debug}

	for _, level := range levels {
		result = append(result, level.ToString())
	}
	return result
}
