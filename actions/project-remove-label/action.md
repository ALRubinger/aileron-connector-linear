+++
name = "project-remove-label"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/project-remove-label@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["project_remove_label"]

[match]
intent = "Removes a label from a project."

[[execute]]
id = "project"
connector = "github://ALRubinger/aileron-connector-linear"
op = "project_remove_label"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the project to remove the label from."
required = true

[[inputs]]
name = "labelId"
type = "String"
description = "The identifier of the label to remove."
required = true
+++
