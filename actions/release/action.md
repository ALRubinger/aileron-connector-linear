+++
name = "release"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/release@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["release"]

[match]
intent = "Fetch a single release by its UUID or slug identifier."

[[execute]]
id = "release"
connector = "github://ALRubinger/aileron-connector-linear"
op = "release"
idempotent = true

[[inputs]]
name = "id"
type = "String"
description = ""
required = true
+++
