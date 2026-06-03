+++
name = "workflow-state-archived"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/workflow-state-archived@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["workflow_state_archived"]

[match]
intent = "Triggered when a workflow state is archived"

[[execute]]
id = "workflow"
connector = "github://ALRubinger/aileron-connector-linear"
op = "workflow_state_archived"
idempotent = true
+++
