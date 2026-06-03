+++
name = "initiative-update"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/initiative-update@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["initiative_update"]

[match]
intent = "Returns a single initiative update by its identifier."

[[execute]]
id = "initiative"
connector = "github://ALRubinger/aileron-connector-linear"
op = "initiative_update"
idempotent = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the initiative update to retrieve."
required = true
+++
