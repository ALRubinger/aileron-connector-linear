+++
name = "project-created"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/project-created@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["project_created"]

[match]
intent = "Triggered when a project is created"

[[execute]]
id = "project"
connector = "github://ALRubinger/aileron-connector-linear"
op = "project_created"
idempotent = true
+++
