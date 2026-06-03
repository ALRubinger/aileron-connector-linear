+++
name = "attachment-sync-to-slack"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/attachment-sync-to-slack@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["attachment_sync_to_slack"]

[match]
intent = "Begin syncing the thread for an existing Slack message attachment with a comment thread on its issue."

[[execute]]
id = "attachment"
connector = "github://ALRubinger/aileron-connector-linear"
op = "attachment_sync_to_slack"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The ID of the Slack attachment to begin syncing."
required = true
+++
