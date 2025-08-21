package main

import (
	"errors"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

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

func TestExtract(t *testing.T) {
	type testCase struct {
		name    string
		errFunc func() error
		wantErr string
	}

	tests := []testCase{
		{
			name: "nil error",
			errFunc: func() error {
				return nil
			},
			wantErr: "",
		},
		{
			name: "plain error without Inject",
			errFunc: func() error {
				return errors.New("some error")
			},
			wantErr: "some error",
		},
		{
			name: "injected error",
			errFunc: func() error {
				return Inject(errors.New("something went wrong"))
			},
			wantErr: "something went wrong",
		},
		{
			name: "injected error with colons in message",
			errFunc: func() error {
				return Inject(errors.New("another: error"))
			},
			wantErr: "another: error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.errFunc()
			gotFile, gotLine, gotErr := Extract(err)

			gotErrStr := ""
			if gotErr != nil {
				gotErrStr = gotErr.Error()
			}

			if tt.name == "plain error without Inject" || err == nil {
				if gotFile != "" || gotLine != 0 || gotErrStr != tt.wantErr {
					t.Errorf("Extract() = (%q, %d, %q), want (%q, %d, %q)",
						gotFile, gotLine, gotErrStr,
						"", 0, tt.wantErr)
				}
			} else {
				if !strings.HasSuffix(gotFile, "error_test.go") || gotLine <= 0 || gotErrStr != tt.wantErr {
					t.Errorf("Extract() = (%q, %d, %q), want (error_test.go, >0, %q)",
						gotFile, gotLine, gotErrStr, tt.wantErr)
				}
			}
		})
	}
}
