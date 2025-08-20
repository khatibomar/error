package main

import (
	"errors"
	"strings"
	"testing"
)

func TestError_Nil(t *testing.T) {
	var err error
	if got := Error(err); got != nil {
		t.Errorf("expected nil, got %v", got)
	}
}

func TestError_AddsPrefix(t *testing.T) {
	err := errors.New("potato")
	wrapped := Error(err)

	msg := wrapped.Error()
	if !strings.Contains(msg, "potato") {
		t.Errorf("expected message to contain 'potato', got %q", msg)
	}
	if !strings.Contains(msg, ".go:") {
		t.Errorf("expected caller prefix, got %q", msg)
	}
}

func TestError_RewrapReplacesPrefix(t *testing.T) {
	err := errors.New("potato")
	first := Error(err)
	second := Error(first)

	firstMsg := first.Error()
	secondMsg := second.Error()

	if !strings.Contains(secondMsg, "potato") {
		t.Errorf("expected message to contain 'potato', got %q", secondMsg)
	}

	if count := strings.Count(secondMsg, ".go:"); count != 1 {
		t.Errorf("expected 1 prefix, got %d in %q", count, secondMsg)
	}

	if firstMsg == secondMsg {
		t.Errorf("expected new prefix, but got same string %q", secondMsg)
	}
}

func TestError_LongMessage(t *testing.T) {
	msg := "this is a very long error message with symbols !@#$%^&*()"
	err := errors.New(msg)
	wrapped := Error(err)

	if !strings.Contains(wrapped.Error(), msg) {
		t.Errorf("expected message to be preserved, got %q", wrapped.Error())
	}
}

func TestError_AlreadyPrefixedEdgeCase(t *testing.T) {
	err := errors.New("something: not a prefix: message")
	wrapped := Error(err)
	eMsg := wrapped.Error()

	if !strings.Contains(eMsg, "something: not a prefix: message") {
		t.Errorf("expected original message preserved, got %q", eMsg)
	}
}
