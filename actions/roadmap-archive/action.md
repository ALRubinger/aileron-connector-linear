+++
name = "roadmap-archive"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/roadmap-archive@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["roadmap_archive"]

[match]
intent = "Archives a roadmap."

[[execute]]
id = "roadmap"
connector = "github://ALRubinger/aileron-connector-linear"
op = "roadmap_archive"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the roadmap to archive."
required = true
+++
