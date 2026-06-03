package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"strings"
	"testing"
)

// errReader is an io.Reader that always fails. Used to exercise the
// stdin-read error branch.
type errReader struct{}

func (errReader) Read([]byte) (int, error) { return 0, errors.New("simulated read failure") }

// TestRun covers the connector's stdin → stdout JSON contract: every
// op declared in the manifest currently returns not_implemented, ops
// not in the manifest return unknown_op, and malformed input returns
// connector_runtime_error. Each of these is part of the contract
// callers (Aileron runtime) rely on.
func TestRun(t *testing.T) {
	t.Run("known op returns not_implemented", func(t *testing.T) {
		for op := range knownOps {
			t.Run(op, func(t *testing.T) {
				stdin := bytes.NewBufferString(`{"op":"` + op + `","args":{}}`)
				var stdout bytes.Buffer
				code := run(stdin, &stdout)
				if code != 0 {
					t.Fatalf("exit code = %d, want 0", code)
				}
				class := decodeErrorClass(t, stdout.Bytes())
				if class != "not_implemented" {
					t.Errorf("error class = %q, want %q", class, "not_implemented")
				}
			})
		}
	})

	t.Run("unknown op returns unknown_op", func(t *testing.T) {
		stdin := bytes.NewBufferString(`{"op":"obliterate_universe","args":{}}`)
		var stdout bytes.Buffer
		code := run(stdin, &stdout)
		if code != 1 {
			t.Fatalf("exit code = %d, want 1", code)
		}
		class := decodeErrorClass(t, stdout.Bytes())
		if class != "unknown_op" {
			t.Errorf("error class = %q, want %q", class, "unknown_op")
		}
	})

	t.Run("stdin read failure returns connector_runtime_error", func(t *testing.T) {
		var stdout bytes.Buffer
		code := run(errReader{}, &stdout)
		if code != 1 {
			t.Fatalf("exit code = %d, want 1", code)
		}
		class := decodeErrorClass(t, stdout.Bytes())
		if class != "connector_runtime_error" {
			t.Errorf("error class = %q, want %q", class, "connector_runtime_error")
		}
	})

	t.Run("malformed JSON returns connector_runtime_error", func(t *testing.T) {
		stdin := bytes.NewBufferString(`{not valid json`)
		var stdout bytes.Buffer
		code := run(stdin, &stdout)
		if code != 1 {
			t.Fatalf("exit code = %d, want 1", code)
		}
		class := decodeErrorClass(t, stdout.Bytes())
		if class != "connector_runtime_error" {
			t.Errorf("error class = %q, want %q", class, "connector_runtime_error")
		}
	})
}

// TestKnownOpsMatchSuiteToml is a cheap consistency check: every op
// dispatched by knownOps should correspond to an `actions/<name>/`
// directory listed in suite.toml. Hand-rolled rather than a real TOML
// parse — we are not validating the file, just guarding against
// rename-drift.
func TestKnownOpsMatchSuiteToml(t *testing.T) {
	// op (snake_case) → expected actions/<dir> as listed in suite.toml.
	expected := map[string]string{
		"issue":          "actions/get-issue",
		"viewer":         "actions/viewer",
		"issue_create":   "actions/issue-create",
		"comment_create": "actions/comment-create",
	}
	if len(expected) != len(knownOps) {
		t.Fatalf("expected/knownOps length mismatch: %d vs %d", len(expected), len(knownOps))
	}
	for op := range knownOps {
		if _, ok := expected[op]; !ok {
			t.Errorf("knownOps has %q but no expected action mapping", op)
		}
	}
}

func decodeErrorClass(t *testing.T, b []byte) string {
	t.Helper()
	var got output
	if err := json.NewDecoder(strings.NewReader(string(b))).Decode(&got); err != nil {
		t.Fatalf("decode output: %v (raw=%q)", err, b)
	}
	if got.Error == nil {
		t.Fatalf("expected error in output, got %s", b)
	}
	return got.Error.Class
}
