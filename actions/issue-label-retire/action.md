+++
name = "issue-label-retire"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/issue-label-retire@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["issue_label_retire"]

[match]
intent = "Retires a label. Retired labels are still visible but cannot be applied to new issues. Existing issues with the label are not affected."

[[execute]]
id = "issue"
connector = "github://ALRubinger/aileron-connector-linear"
op = "issue_label_retire"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the label to retire."
required = true
+++
