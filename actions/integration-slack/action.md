+++
name = "integration-slack"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/integration-slack@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["integration_slack"]

[match]
intent = "Integrates the workspace with Slack."

[[execute]]
id = "integration"
connector = "github://ALRubinger/aileron-connector-linear"
op = "integration_slack"
idempotent = false

[approval]
required = true

[[inputs]]
name = "code"
type = "String"
description = "The Slack OAuth code."
required = true

[[inputs]]
name = "redirectUri"
type = "String"
description = "The Slack OAuth redirect URI."
required = true

[[inputs]]
name = "shouldUseV2Auth"
type = "Boolean"
description = "[DEPRECATED] Whether or not v2 of Slack OAuth should be used. No longer used."
required = false
+++
