+++
name = "integration"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/integration@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["integration"]

[match]
intent = "Retrieves a single integration by its identifier."

[[execute]]
id = "integration"
connector = "github://ALRubinger/aileron-connector-linear"
op = "integration"
idempotent = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the integration to retrieve."
required = true
+++
