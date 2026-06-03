+++
name = "integration-loom"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/integration-loom@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["integration_loom"]

[match]
intent = "Enables Loom integration for the workspace."

[[execute]]
id = "integration"
connector = "github://ALRubinger/aileron-connector-linear"
op = "integration_loom"
idempotent = false

[approval]
required = true
+++
