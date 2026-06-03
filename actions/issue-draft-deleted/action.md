+++
name = "issue-draft-deleted"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/issue-draft-deleted@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["issue_draft_deleted"]

[match]
intent = "Triggered when an issue draft is deleted"

[[execute]]
id = "issue"
connector = "github://ALRubinger/aileron-connector-linear"
op = "issue_draft_deleted"
idempotent = true
+++
