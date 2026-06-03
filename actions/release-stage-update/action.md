+++
name = "release-stage-update"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/release-stage-update@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["release_stage_update"]

[match]
intent = "Updates an existing release stage. Only started-type stages can be edited. Supports updating name, color, position, and frozen status."

[[execute]]
id = "release"
connector = "github://ALRubinger/aileron-connector-linear"
op = "release_stage_update"
idempotent = false

[approval]
required = true

[[inputs]]
name = "color"
type = "String"
description = "The UI color of the stage as a HEX string."
required = false

[[inputs]]
name = "frozen"
type = "Boolean"
description = "Whether this stage is frozen. Only applicable to started stages."
required = false

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the release stage to update."
required = true

[[inputs]]
name = "name"
type = "String"
description = "The name of the stage."
required = false

[[inputs]]
name = "position"
type = "Float"
description = "The position of the stage."
required = false
+++
