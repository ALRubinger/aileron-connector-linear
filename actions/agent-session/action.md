+++
name = "agent-session"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/agent-session@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["agent_session"]

[match]
intent = "A specific agent session."

[[execute]]
id = "agent"
connector = "github://ALRubinger/aileron-connector-linear"
op = "agent_session"
idempotent = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the agent session to retrieve."
required = true
+++
