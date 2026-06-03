+++
name = "integration-front"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/integration-front@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["integration_front"]

[match]
intent = "Integrates the workspace with Front."

[[execute]]
id = "integration"
connector = "github://ALRubinger/aileron-connector-linear"
op = "integration_front"
idempotent = false

[approval]
required = true

[[inputs]]
name = "code"
type = "String"
description = "The Front OAuth code."
required = true

[[inputs]]
name = "redirectUri"
type = "String"
description = "The Front OAuth redirect URI."
required = true
+++
