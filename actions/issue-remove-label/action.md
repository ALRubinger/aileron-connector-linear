+++
name = "issue-remove-label"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/issue-remove-label@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["issue_remove_label"]

[match]
intent = "Removes a label from an issue."

[[execute]]
id = "issue"
connector = "github://ALRubinger/aileron-connector-linear"
op = "issue_remove_label"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the issue to remove the label from."
required = true

[[inputs]]
name = "labelId"
type = "String"
description = "The identifier of the label to remove."
required = true
+++
