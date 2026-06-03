+++
name = "project-relation"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/project-relation@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["project_relation"]

[match]
intent = "Returns a single project relation by its identifier."

[[execute]]
id = "project"
connector = "github://ALRubinger/aileron-connector-linear"
op = "project_relation"
idempotent = true

[[inputs]]
name = "id"
type = "String"
description = ""
required = true
+++
