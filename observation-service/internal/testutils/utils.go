package testutils

import (
	"bytes"
	"io"
	"log"
	"os"
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

// ReadFile reads a file and returns the byte contents
func ReadFile(filepath string) ([]byte, error) {
	// Open file
	fileObj, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer fileObj.Close()
	// Read contents
	byteValue, err := io.ReadAll(fileObj)
	if err != nil {
		return nil, err
	}
	return byteValue, nil
}
