+++
name = "agent-activity-send-queued"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/agent-activity-send-queued@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["agent_activity_send_queued"]

[match]
intent = "[Internal] Immediately sends a queued prompt activity to the agent, bypassing the queue."

[[execute]]
id = "agent"
connector = "github://ALRubinger/aileron-connector-linear"
op = "agent_activity_send_queued"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The ID of the queued agent activity to send."
required = true
+++
