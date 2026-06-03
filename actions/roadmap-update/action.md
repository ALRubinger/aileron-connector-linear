+++
name = "roadmap-update"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/roadmap-update@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["roadmap_update"]

[match]
intent = "Updates a roadmap."

[[execute]]
id = "roadmap"
connector = "github://ALRubinger/aileron-connector-linear"
op = "roadmap_update"
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
description = "The identifier of the roadmap to update."
required = true

[[inputs]]
name = "name"
type = "String"
description = "The name of the roadmap."
required = false

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
