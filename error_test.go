package main

import (
	"errors"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

func TestExtract(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		wantFile string
		wantLine string
	}{
		{
			name:     "nil error",
			err:      nil,
			wantFile: "",
			wantLine: "",
		},
		{
			name:     "error without colon",
			err:      errors.New("some error"),
			wantFile: "",
			wantLine: "",
		},
		{
			name:     "error with single colon",
			err:      errors.New("example.go: some error"),
			wantFile: "",
			wantLine: "",
		},
		{
			name:     "error with file:line prefix",
			err:      errors.New("example.go:123: something went wrong"),
			wantFile: "example.go",
			wantLine: "123",
		},
		{
			name:     "error with extra colons in message",
			err:      errors.New("example.go:45: another: error"),
			wantFile: "example.go",
			wantLine: "45",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFile, gotLine := Extract(tt.err)
			if gotFile != tt.wantFile || gotLine != tt.wantLine {
				t.Errorf("Extract() = (%q, %q), want (%q, %q)", gotFile, gotLine, tt.wantFile, tt.wantLine)
			}
		})
	}
}

func TestInject(t *testing.T) {
	tests := []struct {
		name            string
		err             error
		wantNil         bool
		wantMsgContains string
	}{
		{
			name:    "nil error",
			err:     nil,
			wantNil: true,
		},
		{
			name:            "simple error",
			err:             errors.New("something went wrong"),
			wantNil:         false,
			wantMsgContains: "something went wrong",
		},
		{
			name:            "error with file:line prefix",
			err:             errors.New("example.go:42: original error"),
			wantNil:         false,
			wantMsgContains: "original error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Inject(tt.err)
			if tt.wantNil {
				if got != nil {
					t.Errorf("expected nil, got %v", got)
				}
				return
			}

			if got == nil {
				t.Fatal("expected non-nil error")
			}

			msg := got.Error()
			if !strings.Contains(msg, tt.wantMsgContains) {
				t.Errorf("expected message to contain %q, got %q", tt.wantMsgContains, msg)
			}

			re := regexp.MustCompile(`^[^:]+:\d+:`)
			if !re.MatchString(msg) {
				t.Errorf("expected message to start with file:line:, got %q", msg)
			}

			parts := strings.SplitN(msg, ":", 3)
			file := parts[0]
			if filepath.Base(file) != file {
				t.Errorf("expected file to be basename, got %q", file)
			}
		})
	}
}
