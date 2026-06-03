+++
name = "issue-archive"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/issue-archive@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["issue_archive"]

[match]
intent = "Archives an issue."

[[execute]]
id = "issue"
connector = "github://ALRubinger/aileron-connector-linear"
op = "issue_archive"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the issue to archive."
required = true

[[inputs]]
name = "trash"
type = "Boolean"
description = "Whether to trash the issue."
required = false
+++
