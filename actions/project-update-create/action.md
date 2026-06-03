+++
name = "project-update-create"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/project-update-create@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["project_update_create"]

[match]
intent = "Creates a new project update."

[[execute]]
id = "project"
connector = "github://ALRubinger/aileron-connector-linear"
op = "project_update_create"
idempotent = false

[approval]
required = true

[[inputs]]
name = "body"
type = "String"
description = "The content of the project update in markdown format."
required = false

[[inputs]]
name = "bodyData"
type = "JSON"
description = "[Internal] The content of the project update as a Prosemirror document."
required = false

[[inputs]]
name = "health"
type = "ProjectUpdateHealthType"
description = "The health of the project at the time of the update."
required = false

[[inputs]]
name = "id"
type = "String"
description = "The identifier. If none is provided, the backend will generate one."
required = false

[[inputs]]
name = "isDiffHidden"
type = "Boolean"
description = "Whether the diff between the current update and the previous one should be hidden."
required = false

[[inputs]]
name = "projectId"
type = "String"
description = "The project to associate the project update with."
required = true
+++
