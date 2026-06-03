+++
name = "git-automation-state-update"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/git-automation-state-update@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["git_automation_state_update"]

[match]
intent = "Updates an existing Git automation rule, including its workflow state, target branch, and triggering event."

[[execute]]
id = "git"
connector = "github://ALRubinger/aileron-connector-linear"
op = "git_automation_state_update"
idempotent = false

[approval]
required = true

[[inputs]]
name = "event"
type = "GitAutomationStates"
description = "The event that triggers the automation."
required = false

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the state to update."
required = true

[[inputs]]
name = "stateId"
type = "String"
description = "The associated workflow state."
required = false

[[inputs]]
name = "targetBranchId"
type = "String"
description = "The associated target branch. If null, all branches are targeted."
required = false
+++
