package main

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

// Error wraps the given error with caller information.
//
// It prepends a "file:line" prefix to the error message, indicating
// where Error was called. If the error already has such a prefix,
// it replaces it with the new caller location.
func Error(err error) error {
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
