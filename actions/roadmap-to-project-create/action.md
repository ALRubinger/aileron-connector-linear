+++
name = "roadmap-to-project-create"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/roadmap-to-project-create@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["roadmap_to_project_create"]

[match]
intent = "[Deprecated] Creates a new roadmap-to-project association. Use InitiativeToProject instead."

[[execute]]
id = "roadmap"
connector = "github://ALRubinger/aileron-connector-linear"
op = "roadmap_to_project_create"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier in UUID v4 format. If none is provided, the backend will generate one."
required = false

[[inputs]]
name = "projectId"
type = "String"
description = "The identifier of the project."
required = true

[[inputs]]
name = "roadmapId"
type = "String"
description = "The identifier of the roadmap."
required = true

[[inputs]]
name = "sortOrder"
type = "Float"
description = "The sort order for the project within its workspace."
required = false
+++
