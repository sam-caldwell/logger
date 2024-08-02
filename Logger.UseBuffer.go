package logger

import (
	"io"
)

// UseBuffer - Configure logger to write to io.Writer
func (log *Logger) UseBuffer(w io.Writer) {
	if w == nil {
		panic("cannot use nil writer")
	}
	log.target = w
}
