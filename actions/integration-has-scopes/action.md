+++
name = "integration-has-scopes"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/integration-has-scopes@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["integration_has_scopes"]

[match]
intent = "Checks if the integration has all required scopes."

[[execute]]
id = "integration"
connector = "github://ALRubinger/aileron-connector-linear"
op = "integration_has_scopes"
idempotent = true

[[inputs]]
name = "integrationId"
type = "String"
description = "The integration ID."
required = true

[[inputs]]
name = "scopes"
type = "String"
description = "Required scopes."
required = true
+++
