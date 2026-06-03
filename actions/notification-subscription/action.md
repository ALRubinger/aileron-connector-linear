+++
name = "notification-subscription"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/notification-subscription@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["notification_subscription"]

[match]
intent = "A specific notification subscription by ID."

[[execute]]
id = "notification"
connector = "github://ALRubinger/aileron-connector-linear"
op = "notification_subscription"
idempotent = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the notification subscription to retrieve."
required = true
+++
