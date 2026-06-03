+++
name = "integration-intercom-settings-update"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/integration-intercom-settings-update@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["integration_intercom_settings_update"]

[match]
intent = "[DEPRECATED] Updates settings on the Intercom integration."

[[execute]]
id = "integration"
connector = "github://ALRubinger/aileron-connector-linear"
op = "integration_intercom_settings_update"
idempotent = false

[approval]
required = true

[[inputs]]
name = "automateTicketReopeningOnCancellation"
type = "Boolean"
description = "Whether a ticket should be automatically reopened when its linked Linear issue is canceled."
required = false

[[inputs]]
name = "automateTicketReopeningOnComment"
type = "Boolean"
description = "Whether a ticket should be automatically reopened when a comment is posted on its linked Linear issue"
required = false

[[inputs]]
name = "automateTicketReopeningOnCompletion"
type = "Boolean"
description = "Whether a ticket should be automatically reopened when its linked Linear issue is completed."
required = false

[[inputs]]
name = "automateTicketReopeningOnProjectCancellation"
type = "Boolean"
description = "Whether a ticket should be automatically reopened when its linked Linear project is canceled."
required = false

[[inputs]]
name = "automateTicketReopeningOnProjectCompletion"
type = "Boolean"
description = "Whether a ticket should be automatically reopened when its linked Linear project is completed."
required = false

[[inputs]]
name = "disableCustomerRequestsAutoCreation"
type = "Boolean"
description = "[ALPHA] Whether customer and customer requests should not be automatically created when conversations are linked to a Linear issue."
required = false

[[inputs]]
name = "enableAiIntake"
type = "Boolean"
description = "Whether Linear Agent should be enabled for this integration."
required = false

[[inputs]]
name = "sendNoteOnComment"
type = "Boolean"
description = "Whether an internal message should be added when someone comments on an issue."
required = false

[[inputs]]
name = "sendNoteOnStatusChange"
type = "Boolean"
description = "Whether an internal message should be added when a Linear issue changes status (for status types except completed or canceled)."
required = false
+++
