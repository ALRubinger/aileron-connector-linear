+++
name = "initiative-to-project-update"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/initiative-to-project-update@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["initiative_to_project_update"]

[match]
intent = "Updates an initiative-to-project association, such as its sort order."

[[execute]]
id = "initiative"
connector = "github://ALRubinger/aileron-connector-linear"
op = "initiative_to_project_update"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the initiativeToProject to update."
required = true

[[inputs]]
name = "sortOrder"
type = "Float"
description = "The sort order for the project within the initiative."
required = false
+++
