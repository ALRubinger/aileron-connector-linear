+++
name = "integration-opsgenie-connect"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/integration-opsgenie-connect@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["integration_opsgenie_connect"]

[match]
intent = "[INTERNAL] Integrates the workspace with Opsgenie."

[[execute]]
id = "integration"
connector = "github://ALRubinger/aileron-connector-linear"
op = "integration_opsgenie_connect"
idempotent = false

[approval]
required = true

[[inputs]]
name = "apiKey"
type = "String"
description = "The Opsgenie API key."
required = true
+++
