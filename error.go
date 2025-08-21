package main

import (
	"errors"
	"fmt"
	"path/filepath"
	"runtime"
)

// callerError wraps an existing error with information about the file and line number
// where the error was injected.
type callerError struct {
	file string
	line int
	err  error
}

// Error implements the error interface for callerError.
// It returns a formatted string: "filename:line: original error".
func (e *callerError) Error() string {
	return fmt.Sprintf("%s:%d: %s", e.file, e.line, e.err.Error())
}

// Unwrap allows callerError to support error unwrapping.
// This enables usage with errors.Is and errors.As.
func (e *callerError) Unwrap() error {
	return e.err
}

// Inject wraps the given error with file and line information of the caller.
// If the provided error is nil, Inject returns nil.
//
// Example:
//
//	err := errors.New("something went wrong")
//	err = Inject(err)
//	fmt.Println(err) // might print: "main.go:42: something went wrong"
func Inject(err error) error {
	if err == nil {
		return nil
	}

	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file, line = "???", 0
	}

	return &callerError{
		file: filepath.Base(file),
		line: line,
		err:  err,
	}
}

// Extract retrieves the original error and the file and line information
// from an error injected with Inject. If the error is not a callerError,
// Extract returns the original error and empty file and line values.
//
// Example:
//
//	err := errors.New("something went wrong")
//	injectedErr := Inject(err)
//	file, line, original := Extract(injectedErr)
//	fmt.Println(file, line, original) // might print: "main.go 42 something went wrong"
func Extract(err error) (string, int, error) {
	var ce *callerError
	if errors.As(err, &ce) {
		return ce.file, ce.line, ce.err
	}
	return "", 0, err
}
