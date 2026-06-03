+++
name = "notification-created"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/notification-created@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["notification_created"]

[match]
intent = "Triggered when a notification is created"

[[execute]]
id = "notification"
connector = "github://ALRubinger/aileron-connector-linear"
op = "notification_created"
idempotent = true
+++
