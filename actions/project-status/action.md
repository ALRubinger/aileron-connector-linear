+++
name = "project-status"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/project-status@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["project_status"]

[match]
intent = "Returns a single project status by its identifier."

[[execute]]
id = "project"
connector = "github://ALRubinger/aileron-connector-linear"
op = "project_status"
idempotent = true

[[inputs]]
name = "id"
type = "String"
description = ""
required = true
+++
