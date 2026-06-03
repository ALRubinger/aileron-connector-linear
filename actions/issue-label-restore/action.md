+++
name = "issue-label-restore"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/issue-label-restore@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["issue_label_restore"]

[match]
intent = "Restores a previously retired label, making it available for use again."

[[execute]]
id = "issue"
connector = "github://ALRubinger/aileron-connector-linear"
op = "issue_label_restore"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the label to restore."
required = true
+++
