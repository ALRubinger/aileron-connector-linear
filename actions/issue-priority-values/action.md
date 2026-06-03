+++
name = "issue-priority-values"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/issue-priority-values@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["issue_priority_values"]

[match]
intent = "Issue priority values and corresponding labels."

[[execute]]
id = "issue"
connector = "github://ALRubinger/aileron-connector-linear"
op = "issue_priority_values"
idempotent = true
+++
