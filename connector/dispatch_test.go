package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"strings"
	"testing"
)

// errReader is an io.Reader that always fails. Exercises the
// stdin-read error branch in run().
type errReader struct{}

func (errReader) Read([]byte) (int, error) { return 0, errors.New("simulated read failure") }

// mockTransport captures the request payload and returns a canned
// (body, status, err) tuple. Lets us drive run()/handle() end-to-end
// on the host platform without involving any host import.
type mockTransport struct {
	resp   []byte
	status int
	err    error
	called int
	last   []byte
}

func (m *mockTransport) do(req []byte) ([]byte, int, error) {
	m.called++
	m.last = append([]byte(nil), req...)
	return m.resp, m.status, m.err
}

func TestRun_KnownOpHappyPath(t *testing.T) {
	body := mustJSON(t, map[string]any{
		"data": map[string]any{
			"issue": map[string]any{
				"id":         "iss_1",
				"identifier": "ENG-1",
				"title":      "Demo",
			},
		},
	})
	mt := &mockTransport{resp: body, status: 200}
	stdin := bytes.NewBufferString(`{"op":"issue","args":{"id":"iss_1"}}`)
	var stdout bytes.Buffer

	code := run(stdin, &stdout, mt.do)

	if code != 0 {
		t.Fatalf("exit code = %d, want 0 (stdout=%s)", code, stdout.String())
	}
	if mt.called != 1 {
		t.Errorf("transport calls = %d, want 1", mt.called)
	}
	got := decodeOutput(t, stdout.Bytes())
	if got.Error != nil {
		t.Fatalf("expected success, got error %+v", got.Error)
	}
	issue, ok := got.Output["issue"].(map[string]any)
	if !ok {
		t.Fatalf("output.issue not a map: %+v", got.Output)
	}
	if issue["identifier"] != "ENG-1" {
		t.Errorf("issue.identifier = %v, want ENG-1", issue["identifier"])
	}
}

func TestRun_UnknownOp(t *testing.T) {
	mt := &mockTransport{}
	stdin := bytes.NewBufferString(`{"op":"obliterate_universe","args":{}}`)
	var stdout bytes.Buffer
	code := run(stdin, &stdout, mt.do)
	if code != 1 {
		t.Fatalf("exit code = %d, want 1", code)
	}
	if mt.called != 0 {
		t.Errorf("transport called for unknown op (got %d calls)", mt.called)
	}
	if cls := decodeOutput(t, stdout.Bytes()).Error.Class; cls != "unknown_op" {
		t.Errorf("error class = %q, want unknown_op", cls)
	}
}

func TestRun_MalformedJSON(t *testing.T) {
	mt := &mockTransport{}
	stdin := bytes.NewBufferString(`{not valid json`)
	var stdout bytes.Buffer
	code := run(stdin, &stdout, mt.do)
	if code != 1 {
		t.Fatalf("exit code = %d, want 1", code)
	}
	if cls := decodeOutput(t, stdout.Bytes()).Error.Class; cls != "connector_runtime_error" {
		t.Errorf("error class = %q, want connector_runtime_error", cls)
	}
}

func TestRun_StdinReadFailure(t *testing.T) {
	mt := &mockTransport{}
	var stdout bytes.Buffer
	code := run(errReader{}, &stdout, mt.do)
	if code != 1 {
		t.Fatalf("exit code = %d, want 1", code)
	}
	if cls := decodeOutput(t, stdout.Bytes()).Error.Class; cls != "connector_runtime_error" {
		t.Errorf("error class = %q, want connector_runtime_error", cls)
	}
}

func TestRun_TransportError(t *testing.T) {
	mt := &mockTransport{err: errors.New("simulated transport failure")}
	stdin := bytes.NewBufferString(`{"op":"viewer","args":{}}`)
	var stdout bytes.Buffer
	code := run(stdin, &stdout, mt.do)
	if code != 1 {
		t.Fatalf("exit code = %d, want 1", code)
	}
	out := decodeOutput(t, stdout.Bytes())
	if out.Error == nil {
		t.Fatal("expected error envelope")
	}
	if out.Error.Class != "connector_runtime_error" {
		t.Errorf("error class = %q, want connector_runtime_error", out.Error.Class)
	}
	if !strings.Contains(out.Error.Message, "simulated transport failure") {
		t.Errorf("error message missing inner cause: %q", out.Error.Message)
	}
}

func TestRun_GraphQLErrorResponse(t *testing.T) {
	body := mustJSON(t, map[string]any{
		"data": nil,
		"errors": []any{
			map[string]any{"message": "Entity not found"},
		},
	})
	mt := &mockTransport{resp: body, status: 200}
	stdin := bytes.NewBufferString(`{"op":"issue","args":{"id":"missing"}}`)
	var stdout bytes.Buffer
	code := run(stdin, &stdout, mt.do)
	if code != 1 {
		t.Fatalf("exit code = %d, want 1", code)
	}
	out := decodeOutput(t, stdout.Bytes())
	if out.Error == nil || out.Error.Class != "graphql_error" {
		t.Fatalf("expected graphql_error envelope, got %+v", out.Error)
	}
	if !strings.Contains(out.Error.Message, "Entity not found") {
		t.Errorf("error message missing GraphQL detail: %q", out.Error.Message)
	}
}

func TestRun_HTTPNon2xx(t *testing.T) {
	mt := &mockTransport{resp: []byte(`{"detail":"unauthorized"}`), status: 401}
	stdin := bytes.NewBufferString(`{"op":"viewer","args":{}}`)
	var stdout bytes.Buffer
	code := run(stdin, &stdout, mt.do)
	if code != 1 {
		t.Fatalf("exit code = %d, want 1", code)
	}
	out := decodeOutput(t, stdout.Bytes())
	if out.Error == nil || !strings.Contains(out.Error.Message, "HTTP 401") {
		t.Errorf("expected HTTP 401 in message, got %+v", out.Error)
	}
}

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

func decodeOutput(t *testing.T, b []byte) output {
	t.Helper()
	var got output
	if err := json.NewDecoder(strings.NewReader(string(b))).Decode(&got); err != nil {
		t.Fatalf("decode output: %v (raw=%q)", err, b)
	}
	return got
}

func mustJSON(t *testing.T, v any) []byte {
	t.Helper()
	b, err := json.Marshal(v)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	return b
}
