+++
name = "integration-template-delete"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/integration-template-delete@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["integration_template_delete"]

[match]
intent = "Deletes an integration template connection, removing the link between a template and an integration."

[[execute]]
id = "integration"
connector = "github://ALRubinger/aileron-connector-linear"
op = "integration_template_delete"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the integrationTemplate to delete."
required = true
+++
