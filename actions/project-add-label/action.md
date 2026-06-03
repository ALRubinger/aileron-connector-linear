+++
name = "project-add-label"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/project-add-label@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["project_add_label"]

[match]
intent = "Adds a label to a project."

[[execute]]
id = "project"
connector = "github://ALRubinger/aileron-connector-linear"
op = "project_add_label"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the project to add the label to."
required = true

[[inputs]]
name = "labelId"
type = "String"
description = "The identifier of the label to add."
required = true
+++
