+++
name = "agent-activity-delete-queued"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/agent-activity-delete-queued@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["agent_activity_delete_queued"]

[match]
intent = "[Internal] Deletes a queued prompt activity, removing it from the queue."

[[execute]]
id = "agent"
connector = "github://ALRubinger/aileron-connector-linear"
op = "agent_activity_delete_queued"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The ID of the queued agent activity to delete."
required = true
+++
