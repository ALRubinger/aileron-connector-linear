+++
name = "issue-archived"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/issue-archived@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["issue_archived"]

[match]
intent = "Triggered when an issue is archived"

[[execute]]
id = "issue"
connector = "github://ALRubinger/aileron-connector-linear"
op = "issue_archived"
idempotent = true
+++
