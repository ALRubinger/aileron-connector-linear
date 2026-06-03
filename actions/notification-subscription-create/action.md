+++
name = "notification-subscription-create"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/notification-subscription-create@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["notification_subscription_create"]

[match]
intent = "Creates a new notification subscription for a specific entity. The subscription determines which notification types the authenticated user will receive for the target entity. Exactly one target entity (customer, custom view, cycle, initiative, label, project, team, or user) must be specified."

[[execute]]
id = "notification"
connector = "github://ALRubinger/aileron-connector-linear"
op = "notification_subscription_create"
idempotent = false

[approval]
required = true

[[inputs]]
name = "active"
type = "Boolean"
description = "Whether the subscription is active. Set to false to pause notifications without deleting the subscription."
required = false

[[inputs]]
name = "contextViewType"
type = "ContextViewType"
description = "The type of view to which the notification subscription context is associated with."
required = false

[[inputs]]
name = "customViewId"
type = "String"
description = "The identifier of the custom view to subscribe to."
required = false

[[inputs]]
name = "customerId"
type = "String"
description = "The identifier of the customer to subscribe to."
required = false

[[inputs]]
name = "cycleId"
type = "String"
description = "The identifier of the cycle to subscribe to."
required = false

[[inputs]]
name = "id"
type = "String"
description = "The identifier in UUID v4 format. If none is provided, the backend will generate one."
required = false

[[inputs]]
name = "initiativeId"
type = "String"
description = "The identifier of the initiative to subscribe to."
required = false

[[inputs]]
name = "labelId"
type = "String"
description = "The identifier of the label to subscribe to."
required = false

[[inputs]]
name = "notificationSubscriptionTypes"
type = "String"
description = "The specific notification event types the subscriber wants to receive for the target entity."
required = false

[[inputs]]
name = "projectId"
type = "String"
description = "The identifier of the project to subscribe to."
required = false

[[inputs]]
name = "teamId"
type = "String"
description = "The identifier of the team to subscribe to."
required = false

[[inputs]]
name = "userContextViewType"
type = "UserContextViewType"
description = "The type of user view to which the notification subscription context is associated with."
required = false

[[inputs]]
name = "userId"
type = "String"
description = "The identifier of the user to subscribe to."
required = false
+++
