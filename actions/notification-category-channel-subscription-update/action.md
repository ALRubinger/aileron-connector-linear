+++
name = "notification-category-channel-subscription-update"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/notification-category-channel-subscription-update@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["notification_category_channel_subscription_update"]

[match]
intent = "Subscribes to or unsubscribes from a specific notification category for a given notification channel. For example, subscribe to 'issueAssignment' notifications via the 'email' channel."

[[execute]]
id = "notification"
connector = "github://ALRubinger/aileron-connector-linear"
op = "notification_category_channel_subscription_update"
idempotent = false

[approval]
required = true

[[inputs]]
name = "category"
type = "NotificationCategory"
description = "The notification category to subscribe to or unsubscribe from"
required = true

[[inputs]]
name = "channel"
type = "NotificationChannel"
description = "The notification channel in which to subscribe to or unsubscribe from the category"
required = true

[[inputs]]
name = "subscribe"
type = "Boolean"
description = "True if the user wants to subscribe, false if the user wants to unsubscribe"
required = true
+++
