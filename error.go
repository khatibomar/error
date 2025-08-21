package main

import (
	"errors"
	"fmt"
	"path/filepath"
	"runtime"
)

type callerError struct {
	file string
	line int
	err  error
}

func (e *callerError) Error() string {
	return fmt.Sprintf("%s:%d: %s", e.file, e.line, e.err.Error())
}

func (e *callerError) Unwrap() error {
	return e.err
}

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

func Extract(err error) (string, int, error) {
	var ce *callerError
	if errors.As(err, &ce) {
		return ce.file, ce.line, ce.err
	}
	return "", 0, err
}
