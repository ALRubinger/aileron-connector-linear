+++
name = "project-status-archive"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/project-status-archive@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["project_status_archive"]

[match]
intent = "Archives a project status. The status must not have any active projects assigned to it and must not be the last status of its type."

[[execute]]
id = "project"
connector = "github://ALRubinger/aileron-connector-linear"
op = "project_status_archive"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the project status to archive."
required = true
+++
