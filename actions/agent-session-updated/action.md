+++
name = "agent-session-updated"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/agent-session-updated@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["agent_session_updated"]

[match]
intent = "Triggered when an agent session is updated"

[[execute]]
id = "agent"
connector = "github://ALRubinger/aileron-connector-linear"
op = "agent_session_updated"
idempotent = true
+++
