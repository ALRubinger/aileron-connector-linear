+++
name = "integration-intercom-delete"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/integration-intercom-delete@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["integration_intercom_delete"]

[match]
intent = "Disconnects the workspace from Intercom."

[[execute]]
id = "integration"
connector = "github://ALRubinger/aileron-connector-linear"
op = "integration_intercom_delete"
idempotent = false

[approval]
required = true
+++
