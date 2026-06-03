+++
name = "initiative-to-project"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/initiative-to-project@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["initiative_to_project"]

[match]
intent = "Returns a single initiative-to-project association by its identifier."

[[execute]]
id = "initiative"
connector = "github://ALRubinger/aileron-connector-linear"
op = "initiative_to_project"
idempotent = true

[[inputs]]
name = "id"
type = "String"
description = ""
required = true
+++
