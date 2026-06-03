+++
name = "roadmap-create"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/roadmap-create@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["roadmap_create"]

[match]
intent = "Creates a new roadmap."

[[execute]]
id = "roadmap"
connector = "github://ALRubinger/aileron-connector-linear"
op = "roadmap_create"
idempotent = false

[approval]
required = true

[[inputs]]
name = "color"
type = "String"
description = "The roadmap's color."
required = false

[[inputs]]
name = "description"
type = "String"
description = "The description of the roadmap."
required = false

[[inputs]]
name = "id"
type = "String"
description = "The identifier in UUID v4 format. If none is provided, the backend will generate one."
required = false

[[inputs]]
name = "name"
type = "String"
description = "The name of the roadmap."
required = true

[[inputs]]
name = "ownerId"
type = "String"
description = "The owner of the roadmap."
required = false

[[inputs]]
name = "sortOrder"
type = "Float"
description = "The sort order of the roadmap within the workspace."
required = false
+++
