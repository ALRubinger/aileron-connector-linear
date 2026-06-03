+++
name = "agent-activity-create-prompt"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/agent-activity-create-prompt@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["agent_activity_create_prompt"]

[match]
intent = "[Internal] Creates a prompt agent activity from Linear user input."

[[execute]]
id = "agent"
connector = "github://ALRubinger/aileron-connector-linear"
op = "agent_activity_create_prompt"
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
type = "AgentActivityPromptCreateInputContent"
description = "The content payload of the prompt agent activity."
required = true

[[inputs]]
name = "contextualMetadata"
type = "JSONObject"
description = "[Internal] Metadata about user-provided contextual information for this agent activity."
required = false

[[inputs]]
name = "id"
type = "String"
description = "The identifier in UUID v4 format. If none is provided, the backend will generate one."
required = false

[[inputs]]
name = "queued"
type = "Boolean"
description = "[Internal] If true, the activity is queued and will be sent to the agent when it finishes its current task."
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

[[inputs]]
name = "sourceCommentId"
type = "String"
description = "The comment that contains the content of this activity."
required = false
+++
