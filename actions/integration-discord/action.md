+++
name = "integration-discord"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/integration-discord@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["integration_discord"]

[match]
intent = "Integrates the workspace with Discord."

[[execute]]
id = "integration"
connector = "github://ALRubinger/aileron-connector-linear"
op = "integration_discord"
idempotent = false

[approval]
required = true

[[inputs]]
name = "code"
type = "String"
description = "The Discord OAuth code."
required = true

[[inputs]]
name = "redirectUri"
type = "String"
description = "The Discord OAuth redirect URI."
required = true
+++
