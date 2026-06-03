+++
name = "integration-github-import-connect"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/integration-github-import-connect@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["integration_github_import_connect"]

[match]
intent = "Connects the workspace with the GitHub Import App."

[[execute]]
id = "integration"
connector = "github://ALRubinger/aileron-connector-linear"
op = "integration_github_import_connect"
idempotent = false

[approval]
required = true

[[inputs]]
name = "code"
type = "String"
description = "The GitHub grant code that's exchanged for OAuth tokens."
required = true

[[inputs]]
name = "installationId"
type = "String"
description = "The GitHub data to connect with."
required = true
+++
