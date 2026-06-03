+++
name = "issue-history-created"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/issue-history-created@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["issue_history_created"]

[match]
intent = "Triggered when an issue history is created"

[[execute]]
id = "issue"
connector = "github://ALRubinger/aileron-connector-linear"
op = "issue_history_created"
idempotent = true
+++
