package testutils

import (
	"bytes"
	"log"
)

// CaptureStderrLogs captures standard output and returns it as a string
func CaptureStderrLogs(f func()) string {
	var buf bytes.Buffer

	writer := log.Writer()
	log.SetOutput(&buf)
	f()
	log.SetOutput(writer)

	return buf.String()
}
