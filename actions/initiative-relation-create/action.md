+++
name = "initiative-relation-create"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/initiative-relation-create@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["initiative_relation_create"]

[match]
intent = "Creates a new parent-child relation between two initiatives. The relation cannot create cycles or exceed maximum nesting depth."

[[execute]]
id = "initiative"
connector = "github://ALRubinger/aileron-connector-linear"
op = "initiative_relation_create"
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
description = "The identifier of the parent initiative in the hierarchy."
required = true

[[inputs]]
name = "relatedInitiativeId"
type = "String"
description = "The identifier of the child initiative in the hierarchy."
required = true

[[inputs]]
name = "sortOrder"
type = "Float"
description = "The sort order of the child initiative within its parent."
required = false
+++
