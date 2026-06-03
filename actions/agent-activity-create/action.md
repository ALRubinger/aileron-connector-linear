+++
name = "agent-activity-create"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/agent-activity-create@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["agent_activity_create"]

[match]
intent = "Creates an agent activity."

[[execute]]
id = "agent"
connector = "github://ALRubinger/aileron-connector-linear"
op = "agent_activity_create"
idempotent = false

[approval]
required = true

[[inputs]]
name = "agentSessionId"
type = "String"
description = "The agent session this activity belongs to."
required = true

[[inputs]]
name = "content"
type = "JSONObject"
description = "The content payload of the agent activity. This object is not strictly typed.\nSee https://linear.app/developers/agent-interaction#activity-content-payload for typing details."
required = true

[[inputs]]
name = "contextualMetadata"
type = "JSONObject"
description = "[Internal] Metadata about user-provided contextual information for this agent activity."
required = false

[[inputs]]
name = "ephemeral"
type = "Boolean"
description = "Whether the activity is ephemeral, and should disappear after the next activity. Defaults to false."
required = false

[[inputs]]
name = "id"
type = "String"
description = "The identifier in UUID v4 format. If none is provided, the backend will generate one."
required = false

[[inputs]]
name = "signal"
type = "AgentActivitySignal"
description = "An optional modifier that provides additional instructions on how the activity should be interpreted."
required = false

[[inputs]]
name = "signalMetadata"
type = "JSONObject"
description = "Metadata about this agent activity's signal."
required = false
+++
