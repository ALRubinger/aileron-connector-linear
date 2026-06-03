+++
name = "roadmap-to-project-update"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/roadmap-to-project-update@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["roadmap_to_project_update"]

[match]
intent = "[Deprecated] Updates a roadmap-to-project association. Use InitiativeToProject instead."

[[execute]]
id = "roadmap"
connector = "github://ALRubinger/aileron-connector-linear"
op = "roadmap_to_project_update"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the roadmapToProject to update."
required = true

[[inputs]]
name = "sortOrder"
type = "Float"
description = "The sort order for the project within its workspace."
required = false
+++
