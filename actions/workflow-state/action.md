+++
name = "workflow-state"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/workflow-state@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["workflow_state"]

[match]
intent = "One specific workflow state (issue status), looked up by its unique identifier."

[[execute]]
id = "workflow"
connector = "github://ALRubinger/aileron-connector-linear"
op = "workflow_state"
idempotent = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the workflow state to retrieve."
required = true
+++
