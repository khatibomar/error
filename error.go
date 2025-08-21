package main

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

// Inject wraps the given error with caller information.
//
// It prepends a "file:line" prefix to the error message, indicating
// where Error was called. If the error already has such a prefix,
// it replaces it with the new caller location.
func Inject(err error) error {
	if err == nil {
		return nil
	}

	s := err.Error()

	if i := strings.IndexByte(s, ':'); i != -1 {
		if j := strings.IndexByte(s[i+1:], ':'); j != -1 {
			filePart := strings.TrimSpace(s[:i])
			linePart := strings.TrimSpace(s[i+1 : i+1+j])
			if strings.HasSuffix(filePart, ".go") {
				if _, errConv := strconv.Atoi(linePart); errConv == nil {
					s = strings.TrimSpace(s[i+1+j+1:])
				}
			}
		}
	}

	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file, line = "???", 0
	}

	return fmt.Errorf("%s:%d: %s", filepath.Base(file), line, s)
}

// Extract parses the file and line number from an error string that has a
// `file:line:` prefix.
//
// If the error contains a prefix like "main.go:42: some error", it returns
// the file ("main.go") and line ("42") as strings.
//
// If the error is nil or does not contain a `file:line:` prefix, it returns
// empty strings for both values.
func Extract(err error) (string, string) {
	if err == nil {
		return "", ""
	}

	s := err.Error()

	if i := strings.IndexByte(s, ':'); i != -1 {
		if j := strings.IndexByte(s[i+1:], ':'); j != -1 {
			filePart := strings.TrimSpace(s[:i])
			linePart := strings.TrimSpace(s[i+1 : i+1+j])
			return filePart, linePart
		}
	}

	return "", ""
}
