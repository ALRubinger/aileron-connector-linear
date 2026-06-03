+++
name = "templates"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/templates@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["templates"]

[match]
intent = "All templates in the workspace, including both team-scoped and workspace-level templates."

[[execute]]
id = "templates"
connector = "github://ALRubinger/aileron-connector-linear"
op = "templates"
idempotent = true
+++
