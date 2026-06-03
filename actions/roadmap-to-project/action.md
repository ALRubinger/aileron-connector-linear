+++
name = "roadmap-to-project"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/roadmap-to-project@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["roadmap_to_project"]

[match]
intent = "[Deprecated] Returns a single roadmap-to-project association. Use InitiativeToProject instead."

[[execute]]
id = "roadmap"
connector = "github://ALRubinger/aileron-connector-linear"
op = "roadmap_to_project"
idempotent = true

[[inputs]]
name = "id"
type = "String"
description = ""
required = true
+++
