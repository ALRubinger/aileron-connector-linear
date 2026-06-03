+++
name = "issue-relation-updated"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/issue-relation-updated@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["issue_relation_updated"]

[match]
intent = "Triggered when an issue relation is updated"

[[execute]]
id = "issue"
connector = "github://ALRubinger/aileron-connector-linear"
op = "issue_relation_updated"
idempotent = true
+++
