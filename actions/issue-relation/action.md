+++
name = "issue-relation"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/issue-relation@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["issue_relation"]

[match]
intent = "One specific issue relation, looked up by its unique identifier."

[[execute]]
id = "issue"
connector = "github://ALRubinger/aileron-connector-linear"
op = "issue_relation"
idempotent = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the issue relation to retrieve."
required = true
+++
