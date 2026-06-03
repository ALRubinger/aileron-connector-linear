+++
name = "notifications-unread-count"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/notifications-unread-count@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["notifications_unread_count"]

[match]
intent = "[Internal] A number of unread notifications."

[[execute]]
id = "notifications"
connector = "github://ALRubinger/aileron-connector-linear"
op = "notifications_unread_count"
idempotent = true
+++
