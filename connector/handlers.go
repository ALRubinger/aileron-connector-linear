package main

import "fmt"

// Each handler:
//   1. Validates required args.
//   2. Builds a Linear GraphQL query string + variables map.
//   3. Invokes graphqlCall.
//   4. Returns the `data` subtree directly as the action output, so
//      consuming agents see `{<field>: {...}}` shaped exactly like
//      Linear's own response. No re-shaping or renaming.
//
// Op names match the snake_case `op` fields aileron-codegen emits
// into actions/<name>/action.md; see dispatch.go knownOps.

const issueQuery = `query Issue($id: String!) {
  issue(id: $id) {
    id
    identifier
    title
    description
  }
}`

func handleIssue(args map[string]any, transport Transport) (map[string]any, error) {
	id, err := requireString(args, "id")
	if err != nil {
		return nil, err
	}
	return graphqlCall(transport, issueQuery, map[string]any{"id": id})
}

const viewerQuery = `query Viewer {
  viewer {
    id
    name
    email
  }
}`

func handleViewer(_ map[string]any, transport Transport) (map[string]any, error) {
	return graphqlCall(transport, viewerQuery, nil)
}

const issueCreateMutation = `mutation IssueCreate($input: IssueCreateInput!) {
  issueCreate(input: $input) {
    success
    issue {
      id
      identifier
      title
    }
    lastSyncId
  }
}`

func handleIssueCreate(args map[string]any, transport Transport) (map[string]any, error) {
	input, err := buildIssueCreateInput(args)
	if err != nil {
		return nil, err
	}
	return graphqlCall(transport, issueCreateMutation, map[string]any{"input": input})
}

// buildIssueCreateInput projects args from the action input surface
// (top-level keys: title, teamId, description, assigneeId, priority)
// onto the IssueCreateInput shape Linear expects.
func buildIssueCreateInput(args map[string]any) (map[string]any, error) {
	title, err := requireString(args, "title")
	if err != nil {
		return nil, err
	}
	teamID, err := requireString(args, "teamId")
	if err != nil {
		return nil, err
	}
	out := map[string]any{
		"title":  title,
		"teamId": teamID,
	}
	if v, ok := optionalString(args, "description"); ok {
		out["description"] = v
	}
	if v, ok := optionalString(args, "assigneeId"); ok {
		out["assigneeId"] = v
	}
	if v, ok := optionalInt(args, "priority"); ok {
		out["priority"] = v
	}
	return out, nil
}

const commentCreateMutation = `mutation CommentCreate($input: CommentCreateInput!) {
  commentCreate(input: $input) {
    success
    comment {
      id
      body
    }
    lastSyncId
  }
}`

func handleCommentCreate(args map[string]any, transport Transport) (map[string]any, error) {
	issueID, err := requireString(args, "issueId")
	if err != nil {
		return nil, err
	}
	body, err := requireString(args, "body")
	if err != nil {
		return nil, err
	}
	return graphqlCall(transport, commentCreateMutation, map[string]any{
		"input": map[string]any{
			"issueId": issueID,
			"body":    body,
		},
	})
}

// requireString fetches a string-typed required argument from the
// raw args map. Returns a descriptive error so the agent sees what's
// missing rather than a generic runtime error.
func requireString(args map[string]any, name string) (string, error) {
	v, ok := args[name]
	if !ok {
		return "", fmt.Errorf("missing required arg %q", name)
	}
	s, ok := v.(string)
	if !ok {
		return "", fmt.Errorf("arg %q: expected string, got %T", name, v)
	}
	if s == "" {
		return "", fmt.Errorf("arg %q: must not be empty", name)
	}
	return s, nil
}

// optionalString returns (value, true) when the arg is present and a
// non-empty string. Missing, null, or wrong-typed args yield
// (zero, false); the caller omits them from the GraphQL input.
func optionalString(args map[string]any, name string) (string, bool) {
	v, ok := args[name]
	if !ok || v == nil {
		return "", false
	}
	s, ok := v.(string)
	if !ok || s == "" {
		return "", false
	}
	return s, true
}

// optionalInt returns (value, true) when the arg is present and a
// JSON number. encoding/json decodes numbers into float64 for an
// `any` map, so we accept that and downcast.
func optionalInt(args map[string]any, name string) (int, bool) {
	v, ok := args[name]
	if !ok || v == nil {
		return 0, false
	}
	switch n := v.(type) {
	case int:
		return n, true
	case float64:
		return int(n), true
	}
	return 0, false
}
