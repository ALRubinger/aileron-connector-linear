+++
name = "notification-unarchived"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/notification-unarchived@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["notification_unarchived"]

[match]
intent = "Triggered when a a notification is unarchived"

[[execute]]
id = "notification"
connector = "github://ALRubinger/aileron-connector-linear"
op = "notification_unarchived"
idempotent = true
+++
