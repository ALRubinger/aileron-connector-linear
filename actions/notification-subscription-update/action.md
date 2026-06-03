+++
name = "notification-subscription-update"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/notification-subscription-update@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["notification_subscription_update"]

[match]
intent = "Updates a notification subscription."

[[execute]]
id = "notification"
connector = "github://ALRubinger/aileron-connector-linear"
op = "notification_subscription_update"
idempotent = false

[approval]
required = true

[[inputs]]
name = "active"
type = "Boolean"
description = "Whether the subscription is active."
required = false

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the notification subscription to update."
required = true

[[inputs]]
name = "notificationSubscriptionTypes"
type = "String"
description = "The specific notification event types the subscriber wants to receive. Replaces all previously configured types."
required = false
+++
