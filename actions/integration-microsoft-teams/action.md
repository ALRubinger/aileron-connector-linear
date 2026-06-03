+++
name = "integration-microsoft-teams"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/integration-microsoft-teams@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["integration_microsoft_teams"]

[match]
intent = "Integrates the workspace with Microsoft Teams."

[[execute]]
id = "integration"
connector = "github://ALRubinger/aileron-connector-linear"
op = "integration_microsoft_teams"
idempotent = false

[approval]
required = true

[[inputs]]
name = "code"
type = "String"
description = "The Microsoft OAuth code."
required = true

[[inputs]]
name = "redirectUri"
type = "String"
description = "The Microsoft OAuth redirect URI."
required = true
+++
