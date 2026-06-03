+++
name = "issue-relation-created"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/issue-relation-created@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["issue_relation_created"]

[match]
intent = "Triggered when an issue relation is created"

[[execute]]
id = "issue"
connector = "github://ALRubinger/aileron-connector-linear"
op = "issue_relation_created"
idempotent = true
+++
