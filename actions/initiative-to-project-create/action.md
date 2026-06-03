+++
name = "initiative-to-project-create"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/initiative-to-project-create@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["initiative_to_project_create"]

[match]
intent = "Associates a project with an initiative. A project can only appear once in an initiative hierarchy."

[[execute]]
id = "initiative"
connector = "github://ALRubinger/aileron-connector-linear"
op = "initiative_to_project_create"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier in UUID v4 format. If none is provided, the backend will generate one."
required = false

[[inputs]]
name = "initiativeId"
type = "String"
description = "The identifier of the initiative."
required = true

[[inputs]]
name = "projectId"
type = "String"
description = "The identifier of the project."
required = true

[[inputs]]
name = "sortOrder"
type = "Float"
description = "The sort order for the project within the initiative."
required = false
+++
