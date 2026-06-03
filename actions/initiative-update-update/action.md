+++
name = "initiative-update-update"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/initiative-update-update@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["initiative_update_update"]

[match]
intent = "Updates an initiative update."

[[execute]]
id = "initiative"
connector = "github://ALRubinger/aileron-connector-linear"
op = "initiative_update_update"
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
description = "The content of the update as a Prosemirror document."
required = false

[[inputs]]
name = "health"
type = "InitiativeUpdateHealthType"
description = "The health of the initiative at the time of the update."
required = false

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the initiative update to update."
required = true
+++
