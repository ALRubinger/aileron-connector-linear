+++
name = "initiative-relation"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/initiative-relation@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["initiative_relation"]

[match]
intent = "Returns a single initiative relation by its identifier."

[[execute]]
id = "initiative"
connector = "github://ALRubinger/aileron-connector-linear"
op = "initiative_relation"
idempotent = true

[[inputs]]
name = "id"
type = "String"
description = ""
required = true
+++
