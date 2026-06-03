// Package main is the WASM source for the aileron-connector-linear
// reference connector. It targets Go's native WASI Preview 1
// (`GOOS=wasip1 GOARCH=wasm`) and reads JSON-shaped op dispatches from
// stdin, writing JSON results to stdout.
//
// Build:
//
//	cd connector && GOOS=wasip1 GOARCH=wasm go build -trimpath \
//	  -ldflags="-s -w" -o ../connector.wasm .
//
// Or via Taskfile from the repo root:
//
//	task build
//
// I/O contract (stdin → stdout JSON):
//
//	{"op": "issue", "args": {"id": "<uuid>"}}
//	{"op": "viewer", "args": {}}
//	{"op": "issue_create",
//	 "args": {"title": "...", "teamId": "<uuid>", "description": "...",
//	          "assigneeId": "<uuid>", "priority": 2}}
//	{"op": "comment_create",
//	 "args": {"issueId": "<uuid>", "body": "..."}}
//
// Op names are `snake_case` per the action.md `[[execute]].op` field
// emitted by aileron-codegen; see action.md files under actions/.
//
// This is the scaffolding cut: each handler returns `not_implemented`.
// The real GraphQL POST against api.linear.app and per-op response
// shaping lands in a follow-up PR. See ALRubinger/aileron#915.
package main

import (
	"encoding/json"
	"io"
	"os"
)

type input struct {
	Op   string         `json:"op"`
	Args map[string]any `json:"args"`
}

type output struct {
	Output map[string]any `json:"output,omitempty"`
	Error  *outputError   `json:"error,omitempty"`
}

type outputError struct {
	Class   string `json:"class"`
	Message string `json:"message"`
}

// knownOps are the op names emitted into action.md `[[execute]].op` by
// aileron-codegen from `codegen/gen.yaml`. Keeping the list in lock-step
// with the generated manifests is enforced by `task generate:check` —
// if you add an op below, regenerate `actions/` first.
var knownOps = map[string]bool{
	"issue":          true,
	"viewer":         true,
	"issue_create":   true,
	"comment_create": true,
}

func main() {
	os.Exit(run(os.Stdin, os.Stdout))
}

// run is the testable seam under main(). It reads one JSON op dispatch
// from stdin and writes one JSON result to stdout. Returns the process
// exit code.
func run(stdin io.Reader, stdout io.Writer) int {
	raw, err := io.ReadAll(stdin)
	if err != nil {
		writeError(stdout, "connector_runtime_error", "read_stdin: "+err.Error())
		return 1
	}
	var in input
	if err := json.Unmarshal(raw, &in); err != nil {
		writeError(stdout, "connector_runtime_error", "parse_input: "+err.Error())
		return 1
	}
	if !knownOps[in.Op] {
		writeError(stdout, "unknown_op", "op "+in.Op+" is not exposed by this connector")
		return 1
	}
	writeError(stdout, "not_implemented", "op "+in.Op+" is declared in the manifest but its handler has not landed yet — see ALRubinger/aileron#915")
	return 0
}

func writeError(w io.Writer, class, message string) {
	out := output{Error: &outputError{Class: class, Message: message}}
	if err := json.NewEncoder(w).Encode(out); err != nil {
		os.Stderr.WriteString("aileron-connector-linear: failed to encode error: " + err.Error() + "\n")
	}
}
