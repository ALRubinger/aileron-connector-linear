+++
name = "integration-git-hub-enterprise-server-connect"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/integration-git-hub-enterprise-server-connect@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["integration_git_hub_enterprise_server_connect"]

[match]
intent = "Connects the workspace with a GitHub Enterprise Server."

[[execute]]
id = "integration"
connector = "github://ALRubinger/aileron-connector-linear"
op = "integration_git_hub_enterprise_server_connect"
idempotent = false

[approval]
required = true

[[inputs]]
name = "githubUrl"
type = "String"
description = "The base URL of the GitHub Enterprise Server installation."
required = true

[[inputs]]
name = "organizationName"
type = "String"
description = "The name of GitHub organization."
required = true
+++
