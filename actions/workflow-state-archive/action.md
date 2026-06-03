+++
name = "workflow-state-archive"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/workflow-state-archive@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["workflow_state_archive"]

[match]
intent = "Archives a state. Only states with issues that have all been archived can be archived."

[[execute]]
id = "workflow"
connector = "github://ALRubinger/aileron-connector-linear"
op = "workflow_state_archive"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the state to archive."
required = true
+++
