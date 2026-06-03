+++
name = "project-label-delete"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/project-label-delete@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["project_label_delete"]

[match]
intent = "Deletes a project label."

[[execute]]
id = "project"
connector = "github://ALRubinger/aileron-connector-linear"
op = "project_label_delete"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the label to delete."
required = true
+++
