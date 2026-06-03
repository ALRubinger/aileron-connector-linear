+++
name = "issue-batch-create"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/issue-batch-create@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["issue_batch_create"]

[match]
intent = "Creates a list of issues in one transaction."

[[execute]]
id = "issue"
connector = "github://ALRubinger/aileron-connector-linear"
op = "issue_batch_create"
idempotent = false

[approval]
required = true

[[inputs]]
name = "issues"
type = "IssueCreateInput"
description = "The issues to create."
required = true
+++
