+++
name = "airbyte-integration-connect"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/airbyte-integration-connect@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["airbyte_integration_connect"]

[match]
intent = "Creates an integration api key for Airbyte to connect with Linear."

[[execute]]
id = "airbyte"
connector = "github://ALRubinger/aileron-connector-linear"
op = "airbyte_integration_connect"
idempotent = false

[approval]
required = true

[[inputs]]
name = "apiKey"
type = "String"
description = "Linear export API key."
required = true
+++
