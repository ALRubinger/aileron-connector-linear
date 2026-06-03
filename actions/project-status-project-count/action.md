+++
name = "project-status-project-count"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/project-status-project-count@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["project_status_project_count"]

[match]
intent = "[INTERNAL] Count of projects using this project status across the workspace."

[[execute]]
id = "project"
connector = "github://ALRubinger/aileron-connector-linear"
op = "project_status_project_count"
idempotent = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the project status to find the project count for."
required = true
+++
