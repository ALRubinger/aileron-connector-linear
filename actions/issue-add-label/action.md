+++
name = "issue-add-label"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/issue-add-label@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["issue_add_label"]

[match]
intent = "Adds a label to an issue."

[[execute]]
id = "issue"
connector = "github://ALRubinger/aileron-connector-linear"
op = "issue_add_label"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the issue to add the label to."
required = true

[[inputs]]
name = "labelId"
type = "String"
description = "The identifier of the label to add."
required = true
+++
