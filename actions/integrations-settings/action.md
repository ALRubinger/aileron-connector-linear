+++
name = "integrations-settings"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/integrations-settings@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["integrations_settings"]

[match]
intent = "Retrieves a specific integration settings configuration by its identifier."

[[execute]]
id = "integrations"
connector = "github://ALRubinger/aileron-connector-linear"
op = "integrations_settings"
idempotent = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the integration settings to retrieve."
required = true
+++
