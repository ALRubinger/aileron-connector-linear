+++
name = "workflow-state-update"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/workflow-state-update@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["workflow_state_update"]

[match]
intent = "Updates a state."

[[execute]]
id = "workflow"
connector = "github://ALRubinger/aileron-connector-linear"
op = "workflow_state_update"
idempotent = false

[approval]
required = true

[[inputs]]
name = "color"
type = "String"
description = "The color of the state."
required = false

[[inputs]]
name = "description"
type = "String"
description = "The description of the state."
required = false

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the state to update."
required = true

[[inputs]]
name = "name"
type = "String"
description = "The name of the state."
required = false

[[inputs]]
name = "position"
type = "Float"
description = "The position of the state."
required = false
+++
