+++
name = "integration-template"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/integration-template@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["integration_template"]

[match]
intent = "Retrieves a single integration template connection by its identifier."

[[execute]]
id = "integration"
connector = "github://ALRubinger/aileron-connector-linear"
op = "integration_template"
idempotent = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the integration template to retrieve."
required = true
+++
