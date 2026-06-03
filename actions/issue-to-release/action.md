+++
name = "issue-to-release"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/issue-to-release@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["issue_to_release"]

[match]
intent = "One specific issue-to-release association, looked up by its unique identifier."

[[execute]]
id = "issue"
connector = "github://ALRubinger/aileron-connector-linear"
op = "issue_to_release"
idempotent = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the issue-to-release association to retrieve."
required = true
+++
