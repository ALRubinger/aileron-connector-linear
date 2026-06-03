+++
name = "workflow-state-create"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/workflow-state-create@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["workflow_state_create"]

[match]
intent = "Creates a new state, adding it to the workflow of a team."

[[execute]]
id = "workflow"
connector = "github://ALRubinger/aileron-connector-linear"
op = "workflow_state_create"
idempotent = false

[approval]
required = true

[[inputs]]
name = "color"
type = "String"
description = "The color of the state."
required = true

[[inputs]]
name = "description"
type = "String"
description = "The description of the state."
required = false

[[inputs]]
name = "id"
type = "String"
description = "The identifier in UUID v4 format. If none is provided, the backend will generate one."
required = false

[[inputs]]
name = "name"
type = "String"
description = "The name of the state."
required = true

[[inputs]]
name = "position"
type = "Float"
description = "The position of the state."
required = false

[[inputs]]
name = "teamId"
type = "String"
description = "The team associated with the state."
required = true

[[inputs]]
name = "type"
type = "String"
description = "The workflow state type, which categorizes the state. Valid values: backlog, unstarted, started, completed, canceled. The type determines how the state is treated in workflow progression and reporting."
required = true
+++
