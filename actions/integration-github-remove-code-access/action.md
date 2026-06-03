+++
name = "integration-github-remove-code-access"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/integration-github-remove-code-access@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["integration_github_remove_code_access"]

[match]
intent = "Removes code access from a GitHub integration, downgrading to the basic GitHub App."

[[execute]]
id = "integration"
connector = "github://ALRubinger/aileron-connector-linear"
op = "integration_github_remove_code_access"
idempotent = false

[approval]
required = true

[[inputs]]
name = "integrationId"
type = "String"
description = "The ID of the integration to remove code access from."
required = true
+++
