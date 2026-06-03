+++
name = "integration-asks-connect-channel"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/integration-asks-connect-channel@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["integration_asks_connect_channel"]

[match]
intent = "Connect a Slack channel to Asks."

[[execute]]
id = "integration"
connector = "github://ALRubinger/aileron-connector-linear"
op = "integration_asks_connect_channel"
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
+++
