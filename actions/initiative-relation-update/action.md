+++
name = "initiative-relation-update"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/initiative-relation-update@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["initiative_relation_update"]

[match]
intent = "Updates an initiative relation."

[[execute]]
id = "initiative"
connector = "github://ALRubinger/aileron-connector-linear"
op = "initiative_relation_update"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the initiative relation to update."
required = true

[[inputs]]
name = "sortOrder"
type = "Float"
description = "The sort order of the child initiative within its parent."
required = false
+++
