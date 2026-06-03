+++
name = "push-subscription-create"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/push-subscription-create@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["push_subscription_create"]

[match]
intent = "Creates a push subscription for the authenticated user's current device or browser. If a subscription already exists for the same session, the old one is replaced."

[[execute]]
id = "push"
connector = "github://ALRubinger/aileron-connector-linear"
op = "push_subscription_create"
idempotent = false

[approval]
required = true

[[inputs]]
name = "data"
type = "String"
description = "The push subscription data in stringified JSON format. For web subscriptions, this must contain keys, endpoint, and expirationTime fields per the Web Push API specification. For mobile subscriptions, this contains the device token."
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier in UUID v4 format. If none is provided, the backend will generate one."
required = false

[[inputs]]
name = "type"
type = "PushSubscriptionType"
description = "The type of push subscription: 'web' for browser-based Web Push API, 'apple' for Apple Push Notification service (or 'appleDevelopment' for sandbox), or 'firebase' for Firebase Cloud Messaging (Android)."
required = false
+++
