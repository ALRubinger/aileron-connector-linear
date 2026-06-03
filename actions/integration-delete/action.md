+++
name = "integration-delete"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/integration-delete@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["integration_delete"]

[match]
intent = "Deletes an integration."

[[execute]]
id = "integration"
connector = "github://ALRubinger/aileron-connector-linear"
op = "integration_delete"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the integration to delete."
required = true

[[inputs]]
name = "skipInstallationDeletion"
type = "Boolean"
description = "Whether to skip deleting the installation on the external service side."
required = false
+++
