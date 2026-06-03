+++
name = "triage-responsibility-create"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/triage-responsibility-create@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["triage_responsibility_create"]

[match]
intent = "Creates a new triage responsibility."

[[execute]]
id = "triage"
connector = "github://ALRubinger/aileron-connector-linear"
op = "triage_responsibility_create"
idempotent = false

[approval]
required = true

[[inputs]]
name = "action"
type = "String"
description = "The action to take when an issue is added to triage."
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier in UUID v4 format. If none is provided, the backend will generate one."
required = false

[[inputs]]
name = "manualSelection"
type = "TriageResponsibilityManualSelectionInput"
description = "The manual selection of users responsible for triage."
required = false

[[inputs]]
name = "teamId"
type = "String"
description = "The identifier of the team associated with the triage responsibility."
required = true

[[inputs]]
name = "timeScheduleId"
type = "String"
description = "The identifier of the time schedule used for scheduling triage responsibility"
required = false
+++
