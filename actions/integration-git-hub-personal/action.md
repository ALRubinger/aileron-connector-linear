+++
name = "integration-git-hub-personal"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/integration-git-hub-personal@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["integration_git_hub_personal"]

[match]
intent = "Connect your GitHub account to Linear."

[[execute]]
id = "integration"
connector = "github://ALRubinger/aileron-connector-linear"
op = "integration_git_hub_personal"
idempotent = false

[approval]
required = true

[[inputs]]
name = "code"
type = "String"
description = "The GitHub OAuth code."
required = true

[[inputs]]
name = "codeAccess"
type = "Boolean"
description = "Whether to connect with code access."
required = false

[[inputs]]
name = "enterpriseUrl"
type = "String"
description = "The enterprise URL for GitHub Enterprise Cloud connections. Omit for github.com."
required = false
+++
