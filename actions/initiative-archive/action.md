+++
name = "initiative-archive"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/initiative-archive@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["initiative_archive"]

[match]
intent = "Archives an initiative."

[[execute]]
id = "initiative"
connector = "github://ALRubinger/aileron-connector-linear"
op = "initiative_archive"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the initiative to archive."
required = true
+++
