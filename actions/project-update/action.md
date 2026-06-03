+++
name = "project-update"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/project-update@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["project_update"]

[match]
intent = "Returns a single project update by its identifier."

[[execute]]
id = "project"
connector = "github://ALRubinger/aileron-connector-linear"
op = "project_update"
idempotent = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the project update to retrieve."
required = true
+++
