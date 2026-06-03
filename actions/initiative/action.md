+++
name = "initiative"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/initiative@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["initiative"]

[match]
intent = "Returns a single initiative by its identifier or URL slug."

[[execute]]
id = "initiative"
connector = "github://ALRubinger/aileron-connector-linear"
op = "initiative"
idempotent = true

[[inputs]]
name = "id"
type = "String"
description = ""
required = true
+++
