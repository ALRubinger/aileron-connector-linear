+++
name = "project-reassign-status"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/project-reassign-status@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["project_reassign_status"]

[match]
intent = "[INTERNAL] Reassigns all projects from one project status to another. Used when archiving or deleting a project status."

[[execute]]
id = "project"
connector = "github://ALRubinger/aileron-connector-linear"
op = "project_reassign_status"
idempotent = false

[approval]
required = true

[[inputs]]
name = "newProjectStatusId"
type = "String"
description = "The identifier of the new project status to update the projects to."
required = true

[[inputs]]
name = "originalProjectStatusId"
type = "String"
description = "The identifier of the project status with which projects will be updated."
required = true
+++
