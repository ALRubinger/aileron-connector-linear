+++
name = "issue-unarchived"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/issue-unarchived@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["issue_unarchived"]

[match]
intent = "Triggered when a an issue is unarchived"

[[execute]]
id = "issue"
connector = "github://ALRubinger/aileron-connector-linear"
op = "issue_unarchived"
idempotent = true
+++
