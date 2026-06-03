+++
name = "triage-responsibility-update"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/triage-responsibility-update@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["triage_responsibility_update"]

[match]
intent = "Updates an existing triage responsibility."

[[execute]]
id = "triage"
connector = "github://ALRubinger/aileron-connector-linear"
op = "triage_responsibility_update"
idempotent = false

[approval]
required = true

[[inputs]]
name = "action"
type = "String"
description = "The action to take when an issue is added to triage."
required = false

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the triage responsibility to update."
required = true

[[inputs]]
name = "manualSelection"
type = "TriageResponsibilityManualSelectionInput"
description = "The manual selection of users responsible for triage."
required = false

[[inputs]]
name = "timeScheduleId"
type = "String"
description = "The identifier of the time schedule used for scheduling triage responsibility."
required = false
+++
