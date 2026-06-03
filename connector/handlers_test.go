package main

import (
	"encoding/json"
	"errors"
	"strings"
	"testing"
)

// TestHandlers_GraphQLQueryPayload asserts on the GraphQL request
// payload each handler issues — query string and variables shape.
// Implementation-mirroring is avoided: we assert on the JSON envelope
// the connector emits to the host, which is the contract Linear sees.
func TestHandlers_GraphQLQueryPayload(t *testing.T) {
	cases := []struct {
		name      string
		args      map[string]any
		invoke    func(map[string]any, Transport) (map[string]any, error)
		wantQuery string
		wantVars  map[string]any
	}{
		{
			name:      "issue includes id variable",
			args:      map[string]any{"id": "iss_1"},
			invoke:    handleIssue,
			wantQuery: "issue(id: $id)",
			wantVars:  map[string]any{"id": "iss_1"},
		},
		{
			name:      "viewer takes no variables",
			args:      map[string]any{},
			invoke:    handleViewer,
			wantQuery: "viewer {",
			wantVars:  nil,
		},
		{
			name: "issueCreate flattens optional fields",
			args: map[string]any{
				"title":       "Demo",
				"teamId":      "team_1",
				"description": "Body",
				"priority":    2,
			},
			invoke:    handleIssueCreate,
			wantQuery: "issueCreate(input: $input)",
			wantVars: map[string]any{"input": map[string]any{
				"title":       "Demo",
				"teamId":      "team_1",
				"description": "Body",
				"priority":    2,
			}},
		},
		{
			name: "commentCreate carries issueId + body",
			args: map[string]any{
				"issueId": "iss_1",
				"body":    "lgtm",
			},
			invoke:    handleCommentCreate,
			wantQuery: "commentCreate(input: $input)",
			wantVars: map[string]any{"input": map[string]any{
				"issueId": "iss_1",
				"body":    "lgtm",
			}},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			body := []byte(`{"data":{}}`)
			mt := &mockTransport{resp: body, status: 200}
			if _, err := tc.invoke(tc.args, mt.do); err != nil {
				t.Fatalf("handler: %v", err)
			}
			env := unwrapEnvelope(t, mt.last)
			if !strings.Contains(env.Body.Query, tc.wantQuery) {
				t.Errorf("query missing %q, got %q", tc.wantQuery, env.Body.Query)
			}
			if tc.wantVars == nil {
				if env.Body.Variables != nil {
					t.Errorf("expected no variables, got %+v", env.Body.Variables)
				}
				return
			}
			gotVars := normalizeJSONNumbers(env.Body.Variables).(map[string]any)
			if !equalAny(gotVars, tc.wantVars) {
				t.Errorf("variables mismatch:\n  got:  %+v\n  want: %+v", gotVars, tc.wantVars)
			}
		})
	}
}

func TestHandlers_MissingRequiredArg(t *testing.T) {
	mt := &mockTransport{}
	cases := []struct {
		name   string
		args   map[string]any
		invoke func(map[string]any, Transport) (map[string]any, error)
	}{
		{"issue without id", map[string]any{}, handleIssue},
		{"issueCreate without title", map[string]any{"teamId": "team_1"}, handleIssueCreate},
		{"issueCreate without teamId", map[string]any{"title": "x"}, handleIssueCreate},
		{"commentCreate without issueId", map[string]any{"body": "x"}, handleCommentCreate},
		{"commentCreate without body", map[string]any{"issueId": "x"}, handleCommentCreate},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := tc.invoke(tc.args, mt.do)
			if err == nil {
				t.Fatal("expected error for missing required arg")
			}
			if mt.called != 0 {
				t.Errorf("transport called despite arg-validation failure")
			}
			mt.called = 0
		})
	}
}

func TestHandlers_TransportFailurePropagates(t *testing.T) {
	mt := &mockTransport{err: errors.New("boom")}
	_, err := handleViewer(nil, mt.do)
	if err == nil {
		t.Fatal("expected error from transport failure")
	}
	if !strings.Contains(err.Error(), "boom") {
		t.Errorf("error %q does not contain underlying cause", err.Error())
	}
}

func TestHandlers_GraphQLErrorsSurfaceAsGraphqlErrors(t *testing.T) {
	body := []byte(`{"errors":[{"message":"unauthorized"}]}`)
	mt := &mockTransport{resp: body, status: 200}
	_, err := handleViewer(nil, mt.do)
	if err == nil {
		t.Fatal("expected error from graphql errors[]")
	}
	var gerr *graphqlErrors
	if !errors.As(err, &gerr) {
		t.Fatalf("expected *graphqlErrors, got %T (%v)", err, err)
	}
	if len(gerr.Errors) != 1 || gerr.Errors[0].Message != "unauthorized" {
		t.Errorf("unexpected graphqlErrors content: %+v", gerr.Errors)
	}
}

// envelope is the host-import HTTP envelope shape graphqlCall builds.
// Only the fields the tests care about are decoded.
type envelope struct {
	Method     string            `json:"method"`
	URL        string            `json:"url"`
	Credential string            `json:"credential"`
	Headers    map[string]string `json:"headers"`
	Body       graphqlPayload    `json:"-"`
	RawBody    string            `json:"body"`
}

func unwrapEnvelope(t *testing.T, raw []byte) envelope {
	t.Helper()
	var env envelope
	if err := json.Unmarshal(raw, &env); err != nil {
		t.Fatalf("decode envelope: %v (raw=%q)", err, raw)
	}
	if env.Credential != "api_key" {
		t.Errorf("envelope.credential = %q, want api_key", env.Credential)
	}
	if env.URL != graphqlEndpoint {
		t.Errorf("envelope.url = %q, want %q", env.URL, graphqlEndpoint)
	}
	if env.Method != "POST" {
		t.Errorf("envelope.method = %q, want POST", env.Method)
	}
	if err := json.Unmarshal([]byte(env.RawBody), &env.Body); err != nil {
		t.Fatalf("decode graphql payload: %v (body=%q)", err, env.RawBody)
	}
	return env
}

// normalizeJSONNumbers recursively converts json-decoded numbers
// (float64) into int where they round-trip cleanly. Lets tests
// compare {"priority": 2} against an int-typed expectation.
func normalizeJSONNumbers(v any) any {
	switch t := v.(type) {
	case map[string]any:
		out := make(map[string]any, len(t))
		for k, x := range t {
			out[k] = normalizeJSONNumbers(x)
		}
		return out
	case []any:
		out := make([]any, len(t))
		for i, x := range t {
			out[i] = normalizeJSONNumbers(x)
		}
		return out
	case float64:
		if t == float64(int(t)) {
			return int(t)
		}
		return t
	}
	return v
}

func equalAny(a, b any) bool {
	switch av := a.(type) {
	case map[string]any:
		bv, ok := b.(map[string]any)
		if !ok || len(av) != len(bv) {
			return false
		}
		for k, x := range av {
			if !equalAny(x, bv[k]) {
				return false
			}
		}
		return true
	case []any:
		bv, ok := b.([]any)
		if !ok || len(av) != len(bv) {
			return false
		}
		for i, x := range av {
			if !equalAny(x, bv[i]) {
				return false
			}
		}
		return true
	default:
		return a == b
	}
}
