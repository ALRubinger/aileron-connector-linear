+++
name = "issue-label"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/issue-label@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["issue_label"]

[match]
intent = "One specific label, looked up by its unique identifier."

[[execute]]
id = "issue"
connector = "github://ALRubinger/aileron-connector-linear"
op = "issue_label"
idempotent = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the label to retrieve."
required = true
+++
