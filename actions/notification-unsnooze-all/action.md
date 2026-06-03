+++
name = "notification-unsnooze-all"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/notification-unsnooze-all@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["notification_unsnooze_all"]

[match]
intent = "Unsnoozes a notification and all related notifications."

[[execute]]
id = "notification"
connector = "github://ALRubinger/aileron-connector-linear"
op = "notification_unsnooze_all"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The id of the notification."
required = false

[[inputs]]
name = "initiativeId"
type = "String"
description = "The id of the initiative related to the notification."
required = false

[[inputs]]
name = "initiativeUpdateId"
type = "String"
description = "The id of the initiative update related to the notification."
required = false

[[inputs]]
name = "issueId"
type = "String"
description = "The id of the issue related to the notification."
required = false

[[inputs]]
name = "oauthClientApprovalId"
type = "String"
description = "The id of the OAuth client approval related to the notification."
required = false

[[inputs]]
name = "projectId"
type = "String"
description = "[DEPRECATED] The id of the project related to the notification."
required = false

[[inputs]]
name = "projectUpdateId"
type = "String"
description = "The id of the project update related to the notification."
required = false

[[inputs]]
name = "unsnoozedAt"
type = "DateTime"
description = "The time when the notification was unsnoozed."
required = true
+++
