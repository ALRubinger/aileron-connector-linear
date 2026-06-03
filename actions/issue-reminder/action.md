+++
name = "issue-reminder"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/issue-reminder@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["issue_reminder"]

[match]
intent = "Adds an issue reminder. Will cause a notification to be sent when the issue reminder time is reached."

[[execute]]
id = "issue"
connector = "github://ALRubinger/aileron-connector-linear"
op = "issue_reminder"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the issue to add a reminder for."
required = true

[[inputs]]
name = "reminderAt"
type = "DateTime"
description = "The time when a reminder notification will be sent."
required = true
+++
