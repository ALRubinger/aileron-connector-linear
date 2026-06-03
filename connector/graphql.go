package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

// graphqlPayload is the JSON shape POSTed to api.linear.app/graphql.
type graphqlPayload struct {
	Query     string         `json:"query"`
	Variables map[string]any `json:"variables,omitempty"`
}

// graphqlResponse is the JSON shape Linear returns.
type graphqlResponse struct {
	Data   map[string]any   `json:"data,omitempty"`
	Errors []graphqlError   `json:"errors,omitempty"`
}

// graphqlError matches Linear's per-error shape (subset of the spec).
type graphqlError struct {
	Message string `json:"message"`
	Path    []any  `json:"path,omitempty"`
}

// graphqlErrors wraps a non-empty errors[] response so callers can
// `errors.As(...)` to detect a GraphQL-level failure separately from
// transport / HTTP-level failures.
type graphqlErrors struct {
	Errors []graphqlError
}

func (e *graphqlErrors) Error() string {
	msgs := make([]string, 0, len(e.Errors))
	for _, gqlErr := range e.Errors {
		msgs = append(msgs, gqlErr.Message)
	}
	return "graphql: " + strings.Join(msgs, "; ")
}

// graphqlCall builds the Aileron-host HTTP request envelope, POSTs it
// via transport, and unmarshals the GraphQL response. Returns the
// parsed `data` map on success; returns a *graphqlErrors if the
// response carries non-empty errors[]; otherwise an error describing
// the transport or HTTP failure.
//
// The `credential: "api_key"` field tells the runtime to inject the
// `Authorization` header host-side per ADR-0005; the connector binary
// never sees the bound key.
func graphqlCall(transport Transport, query string, variables map[string]any) (map[string]any, error) {
	body, err := json.Marshal(graphqlPayload{Query: query, Variables: variables})
	if err != nil {
		return nil, fmt.Errorf("marshal graphql payload: %w", err)
	}
	req, err := json.Marshal(map[string]any{
		"method": "POST",
		"url":    linearGraphQLEndpoint,
		"headers": map[string]string{
			"Accept":       "application/json",
			"Content-Type": "application/json",
		},
		"body":       string(body),
		"credential": "api_key",
	})
	if err != nil {
		return nil, fmt.Errorf("marshal http envelope: %w", err)
	}
	respBody, status, err := transport(req)
	if err != nil {
		return nil, fmt.Errorf("transport: %w", err)
	}
	if status < 200 || status >= 300 {
		return nil, fmt.Errorf("linear graphql HTTP %d: %s", status, truncate(string(respBody), 512))
	}
	var resp graphqlResponse
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return nil, fmt.Errorf("parse graphql response: %w (body=%q)", err, truncate(string(respBody), 512))
	}
	if len(resp.Errors) > 0 {
		return nil, &graphqlErrors{Errors: resp.Errors}
	}
	return resp.Data, nil
}

// truncate caps long strings in error messages so they stay readable
// in the JSON envelope written to stdout.
func truncate(s string, max int) string {
	if len(s) <= max {
		return s
	}
	return s[:max] + "..."
}
