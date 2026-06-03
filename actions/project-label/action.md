+++
name = "project-label"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/project-label@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["project_label"]

[match]
intent = "Returns a single project label by its identifier."

[[execute]]
id = "project"
connector = "github://ALRubinger/aileron-connector-linear"
op = "project_label"
idempotent = true

[[inputs]]
name = "id"
type = "String"
description = ""
required = true
+++
