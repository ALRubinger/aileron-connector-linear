+++
name = "project-status-create"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/project-status-create@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["project_status_create"]

[match]
intent = "Creates a new project status."

[[execute]]
id = "project"
connector = "github://ALRubinger/aileron-connector-linear"
op = "project_status_create"
idempotent = false

[approval]
required = true

[[inputs]]
name = "color"
type = "String"
description = "The UI color of the status as a HEX string."
required = true

[[inputs]]
name = "description"
type = "String"
description = "Description of the status."
required = false

[[inputs]]
name = "id"
type = "String"
description = "The identifier in UUID v4 format. If none is provided, the backend will generate one."
required = false

[[inputs]]
name = "indefinite"
type = "Boolean"
description = "Whether or not a project can be in this status indefinitely."
required = false

[[inputs]]
name = "name"
type = "String"
description = "The name of the status."
required = true

[[inputs]]
name = "position"
type = "Float"
description = "The position of the status in the workspace's project flow."
required = true

[[inputs]]
name = "type"
type = "ProjectStatusType"
description = "The type of the project status."
required = true
+++
