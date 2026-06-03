+++
name = "git-automation-state-create"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/git-automation-state-create@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["git_automation_state_create"]

[match]
intent = "Creates a new Git automation rule that maps a Git event to a workflow state transition for a team."

[[execute]]
id = "git"
connector = "github://ALRubinger/aileron-connector-linear"
op = "git_automation_state_create"
idempotent = false

[approval]
required = true

[[inputs]]
name = "event"
type = "GitAutomationStates"
description = "The event that triggers the automation."
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier in UUID v4 format. If none is provided, the backend will generate one."
required = false

[[inputs]]
name = "stateId"
type = "String"
description = "The associated workflow state. If null, will override default behaviour and take no action."
required = false

[[inputs]]
name = "targetBranchId"
type = "String"
description = "The associated target branch. If null, all branches are targeted."
required = false

[[inputs]]
name = "teamId"
type = "String"
description = "The team associated with the automation state."
required = true
+++
