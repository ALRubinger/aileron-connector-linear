+++
name = "project-update-archive"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/project-update-archive@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["project_update_archive"]

[match]
intent = "Archives a project update."

[[execute]]
id = "project"
connector = "github://ALRubinger/aileron-connector-linear"
op = "project_update_archive"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the project update to archive."
required = true
+++
