+++
name = "notification-update"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/notification-update@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["notification_update"]

[match]
intent = "Updates a notification."

[[execute]]
id = "notification"
connector = "github://ALRubinger/aileron-connector-linear"
op = "notification_update"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the notification to update."
required = true

[[inputs]]
name = "initiativeUpdateId"
type = "String"
description = "The id of the initiative update related to the notification."
required = false

[[inputs]]
name = "projectUpdateId"
type = "String"
description = "The id of the project update related to the notification."
required = false

[[inputs]]
name = "readAt"
type = "DateTime"
description = "The time when notification was marked as read."
required = false

[[inputs]]
name = "snoozedUntilAt"
type = "DateTime"
description = "The time until a notification will be snoozed. After that it will appear in the inbox again."
required = false
+++
