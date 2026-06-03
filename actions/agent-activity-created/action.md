+++
name = "agent-activity-created"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/agent-activity-created@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["agent_activity_created"]

[match]
intent = "Triggered when an agent activity is created"

[[execute]]
id = "agent"
connector = "github://ALRubinger/aileron-connector-linear"
op = "agent_activity_created"
idempotent = true
+++
