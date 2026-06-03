+++
name = "cycle-update"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/cycle-update@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["cycle_update"]

[match]
intent = "Updates a cycle."

[[execute]]
id = "cycle"
connector = "github://ALRubinger/aileron-connector-linear"
op = "cycle_update"
idempotent = false

[approval]
required = true

[[inputs]]
name = "completedAt"
type = "DateTime"
description = "The completion time of the cycle. If null, the cycle hasn't been completed."
required = false

[[inputs]]
name = "description"
type = "String"
description = "The description of the cycle."
required = false

[[inputs]]
name = "endsAt"
type = "DateTime"
description = "The end time of the cycle."
required = false

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the cycle to update."
required = true

[[inputs]]
name = "name"
type = "String"
description = "The custom name of the cycle."
required = false

[[inputs]]
name = "startsAt"
type = "DateTime"
description = "The start time of the cycle."
required = false
+++
