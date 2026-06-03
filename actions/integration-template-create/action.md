+++
name = "integration-template-create"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/integration-template-create@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["integration_template_create"]

[match]
intent = "Creates a new connection between a template and an integration, optionally scoped to a specific external resource such as a Slack channel."

[[execute]]
id = "integration"
connector = "github://ALRubinger/aileron-connector-linear"
op = "integration_template_create"
idempotent = false

[approval]
required = true

[[inputs]]
name = "foreignEntityId"
type = "String"
description = "The foreign identifier in the other service."
required = false

[[inputs]]
name = "id"
type = "String"
description = "The identifier in UUID v4 format. If none is provided, the backend will generate one."
required = false

[[inputs]]
name = "integrationId"
type = "String"
description = "The identifier of the integration."
required = true

[[inputs]]
name = "templateId"
type = "String"
description = "The identifier of the template."
required = true
+++
