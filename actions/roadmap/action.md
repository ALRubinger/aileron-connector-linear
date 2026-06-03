+++
name = "roadmap"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/roadmap@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["roadmap"]

[match]
intent = "[Deprecated] Returns a single roadmap by its identifier. Use initiatives instead."

[[execute]]
id = "roadmap"
connector = "github://ALRubinger/aileron-connector-linear"
op = "roadmap"
idempotent = true

[[inputs]]
name = "id"
type = "String"
description = ""
required = true
+++
