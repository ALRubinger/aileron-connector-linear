+++
name = "project-relation-delete"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/project-relation-delete@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["project_relation_delete"]

[match]
intent = "Deletes a project relation."

[[execute]]
id = "project"
connector = "github://ALRubinger/aileron-connector-linear"
op = "project_relation_delete"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the project relation to delete."
required = true
+++
