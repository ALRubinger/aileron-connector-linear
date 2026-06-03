+++
name = "initiative-to-project-delete"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/initiative-to-project-delete@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["initiative_to_project_delete"]

[match]
intent = "Removes a project from an initiative."

[[execute]]
id = "initiative"
connector = "github://ALRubinger/aileron-connector-linear"
op = "initiative_to_project_delete"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the initiativeToProject to delete."
required = true
+++
