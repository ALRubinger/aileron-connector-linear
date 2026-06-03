+++
name = "initiative-update-create"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/initiative-update-create@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["initiative_update_create"]

[match]
intent = "Creates an initiative update."

[[execute]]
id = "initiative"
connector = "github://ALRubinger/aileron-connector-linear"
op = "initiative_update_create"
idempotent = false

[approval]
required = true

[[inputs]]
name = "body"
type = "String"
description = "The content of the update in markdown format."
required = false

[[inputs]]
name = "bodyData"
type = "JSON"
description = "[Internal] The content of the update as a Prosemirror document."
required = false

[[inputs]]
name = "health"
type = "InitiativeUpdateHealthType"
description = "The health of the initiative at the time of the update."
required = false

[[inputs]]
name = "id"
type = "String"
description = "The identifier in UUID v4 format. If none is provided, the backend will generate one."
required = false

[[inputs]]
name = "initiativeId"
type = "String"
description = "The initiative to associate the update with."
required = true

[[inputs]]
name = "isDiffHidden"
type = "Boolean"
description = "Whether the diff between the current update and the previous one should be hidden."
required = false
+++
