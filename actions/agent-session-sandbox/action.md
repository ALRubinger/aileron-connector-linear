+++
name = "agent-session-sandbox"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/agent-session-sandbox@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["agent_session_sandbox"]

[match]
intent = "[Internal] Retrieves coding agent sandbox details for a given agent session ID."

[[execute]]
id = "agent"
connector = "github://ALRubinger/aileron-connector-linear"
op = "agent_session_sandbox"
idempotent = true

[[inputs]]
name = "agentSessionId"
type = "String"
description = "The identifier of the agent session."
required = true
+++
