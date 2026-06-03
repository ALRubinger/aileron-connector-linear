+++
name = "notification-deleted"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/notification-deleted@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["notification_deleted"]

[match]
intent = "Triggered when a notification is deleted"

[[execute]]
id = "notification"
connector = "github://ALRubinger/aileron-connector-linear"
op = "notification_deleted"
idempotent = true
+++
