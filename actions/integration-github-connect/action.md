+++
name = "integration-github-connect"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/integration-github-connect@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["integration_github_connect"]

[match]
intent = "Connects the workspace with the GitHub App."

[[execute]]
id = "integration"
connector = "github://ALRubinger/aileron-connector-linear"
op = "integration_github_connect"
idempotent = false

[approval]
required = true

[[inputs]]
name = "code"
type = "String"
description = "The GitHub grant code that's exchanged for OAuth tokens."
required = true

[[inputs]]
name = "codeAccess"
type = "Boolean"
description = "Whether the integration should have code access"
required = false

[[inputs]]
name = "confirmReplace"
type = "Boolean"
description = "Set to true to force-replace the existing GitHub App installation (e.g. to ignore repository access conflicts)."
required = false

[[inputs]]
name = "githubHost"
type = "String"
description = "The GitHub Enterprise Cloud host (e.g., 'linear-development.ghe.com'). When provided, this creates a GEC integration."
required = false

[[inputs]]
name = "installationId"
type = "String"
description = "The GitHub data to connect with."
required = true
+++
