+++
name = "notification-archive"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/notification-archive@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["notification_archive"]

[match]
intent = "Archives a notification."

[[execute]]
id = "notification"
connector = "github://ALRubinger/aileron-connector-linear"
op = "notification_archive"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The id of the notification to archive."
required = true
+++
