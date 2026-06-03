+++
name = "project"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/project@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["project"]

[match]
intent = "Returns a single project by its identifier or URL slug."

[[execute]]
id = "project"
connector = "github://ALRubinger/aileron-connector-linear"
op = "project"
idempotent = true

[[inputs]]
name = "id"
type = "String"
description = ""
required = true
+++
