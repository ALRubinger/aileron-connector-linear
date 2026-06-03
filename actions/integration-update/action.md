+++
name = "integration-update"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/integration-update@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["integration_update"]

[match]
intent = "[INTERNAL] Updates the integration."

[[execute]]
id = "integration"
connector = "github://ALRubinger/aileron-connector-linear"
op = "integration_update"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the integration to update."
required = true

[[inputs]]
name = "settings"
type = "IntegrationSettingsInput"
description = "The settings to update."
required = false

[[inputs]]
name = "workflowDefinitionDraftId"
type = "String"
description = "The ID of the workflow definition draft that this integration should be assigned to."
required = false
+++
