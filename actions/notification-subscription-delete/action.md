+++
name = "notification-subscription-delete"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/notification-subscription-delete@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["notification_subscription_delete"]

[match]
intent = "Deletes a notification subscription reference."

[[execute]]
id = "notification"
connector = "github://ALRubinger/aileron-connector-linear"
op = "notification_subscription_delete"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the notification subscription reference to delete."
required = true
+++
