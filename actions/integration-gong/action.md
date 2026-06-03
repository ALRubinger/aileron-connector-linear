+++
name = "integration-gong"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/integration-gong@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["integration_gong"]

[match]
intent = "Integrates the workspace with Gong."

[[execute]]
id = "integration"
connector = "github://ALRubinger/aileron-connector-linear"
op = "integration_gong"
idempotent = false

[approval]
required = true

[[inputs]]
name = "code"
type = "String"
description = "The Gong OAuth code."
required = true

[[inputs]]
name = "redirectUri"
type = "String"
description = "The Gong OAuth redirect URI."
required = true
+++
