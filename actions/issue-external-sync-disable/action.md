+++
name = "issue-external-sync-disable"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/issue-external-sync-disable@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["issue_external_sync_disable"]

[match]
intent = "Disables external sync on an issue."

[[execute]]
id = "issue"
connector = "github://ALRubinger/aileron-connector-linear"
op = "issue_external_sync_disable"
idempotent = false

[approval]
required = true

[[inputs]]
name = "attachmentId"
type = "String"
description = "The ID of the sync attachment to disable."
required = true
+++
