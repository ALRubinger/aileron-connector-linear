+++
name = "project-unarchived"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/project-unarchived@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["project_unarchived"]

[match]
intent = "Triggered when a a project is unarchived"

[[execute]]
id = "project"
connector = "github://ALRubinger/aileron-connector-linear"
op = "project_unarchived"
idempotent = true
+++
