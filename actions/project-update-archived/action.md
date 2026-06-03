+++
name = "project-update-archived"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/project-update-archived@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["project_update_archived"]

[match]
intent = "Triggered when a project update is archived"

[[execute]]
id = "project"
connector = "github://ALRubinger/aileron-connector-linear"
op = "project_update_archived"
idempotent = true
+++
