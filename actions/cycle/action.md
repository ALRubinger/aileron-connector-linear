+++
name = "cycle"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/cycle@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["cycle"]

[match]
intent = "One specific cycle, looked up by ID or slug."

[[execute]]
id = "cycle"
connector = "github://ALRubinger/aileron-connector-linear"
op = "cycle"
idempotent = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the cycle to retrieve."
required = true
+++
