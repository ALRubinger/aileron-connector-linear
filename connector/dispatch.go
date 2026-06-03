// dispatch.go is intentionally untagged so it builds on every Go
// target (host + wasip1). The wasip1-only host-import shim lives in
// main.go; keeping the dispatch core here lets `go test` exercise
// every handler on the host platform with a mock Transport.

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// linearGraphQLEndpoint is the only outbound URL this connector dials.
// The same host:port pair is allowlisted in connector/manifest.toml.
const linearGraphQLEndpoint = "https://api.linear.app/graphql"

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

// Transport sends one Aileron-host HTTP request envelope and returns
// (body, status, err). The wasip1 build provides the real
// implementation via host imports (see main.go); host builds use the
// always-fail stub in transport_host.go or an injected mock in tests.
type Transport func(payload []byte) ([]byte, int, error)

// knownOps must stay in lock-step with the `op` fields emitted into
// actions/<name>/action.md by aileron-codegen. The CI `generate:check`
// step keeps the action manifests in sync with `codegen/gen.yaml`;
// this map then has to track those op names. The dispatch table in
// handle() is the second consistency point.
var knownOps = map[string]bool{
	"issue":          true,
	"viewer":         true,
	"issue_create":   true,
	"comment_create": true,
}

// run is the testable seam under main(). It reads one JSON op
// dispatch from stdin, drives the chosen handler against the supplied
// Transport, and writes one JSON result envelope to stdout.
func run(stdin io.Reader, stdout io.Writer, transport Transport) int {
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
	data, err := handle(in.Op, in.Args, transport)
	if err != nil {
		writeError(stdout, classify(err), err.Error())
		return 1
	}
	writeSuccess(stdout, data)
	return 0
}

// handle dispatches one op to its handler. The caller must have
// already validated that op is in knownOps.
func handle(op string, args map[string]any, transport Transport) (map[string]any, error) {
	switch op {
	case "issue":
		return handleIssue(args, transport)
	case "viewer":
		return handleViewer(args, transport)
	case "issue_create":
		return handleIssueCreate(args, transport)
	case "comment_create":
		return handleCommentCreate(args, transport)
	}
	// Unreachable when knownOps and the switch above stay in sync; the
	// CI tests catch drift between them.
	return nil, fmt.Errorf("unhandled op %q", op)
}

func writeError(w io.Writer, class, message string) {
	out := output{Error: &outputError{Class: class, Message: message}}
	_ = json.NewEncoder(w).Encode(out)
}

func writeSuccess(w io.Writer, payload map[string]any) {
	out := output{Output: payload}
	_ = json.NewEncoder(w).Encode(out)
}

// classify maps a handler error into the Aileron error class string.
// Today every handler error funnels to `connector_runtime_error`; the
// taxonomy expands as specific error paths (rate-limit, not-found,
// validation) call for their own classes.
func classify(err error) string {
	var graphqlErr *graphqlErrors
	if errors.As(err, &graphqlErr) {
		return "graphql_error"
	}
	return "connector_runtime_error"
}
