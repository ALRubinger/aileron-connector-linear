+++
name = "integration-figma"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/integration-figma@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["integration_figma"]

[match]
intent = "Integrates the workspace with Figma."

[[execute]]
id = "integration"
connector = "github://ALRubinger/aileron-connector-linear"
op = "integration_figma"
idempotent = false

[approval]
required = true

[[inputs]]
name = "code"
type = "String"
description = "The Figma OAuth code."
required = true

[[inputs]]
name = "redirectUri"
type = "String"
description = "The Figma OAuth redirect URI."
required = true
+++
