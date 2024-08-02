package logger

// ToString - Print a string for the verbosity value.
func (v *Verbosity) ToString() string {
	switch *v {
	case Normal:
		return "normal"
	case Quiet:
		return "quiet"
	default:
		panic("invalid log verbosity parameter")
	}
}
