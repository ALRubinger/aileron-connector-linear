+++
name = "release-stage"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/release-stage@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["release_stage"]

[match]
intent = "Fetch a single release stage by its UUID."

[[execute]]
id = "release"
connector = "github://ALRubinger/aileron-connector-linear"
op = "release_stage"
idempotent = true

[[inputs]]
name = "id"
type = "String"
description = ""
required = true
+++
