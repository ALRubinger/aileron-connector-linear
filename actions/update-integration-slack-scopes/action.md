+++
name = "update-integration-slack-scopes"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/update-integration-slack-scopes@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["update_integration_slack_scopes"]

[match]
intent = "[Internal] Updates existing Slack and Asks integration scopes."

[[execute]]
id = "update"
connector = "github://ALRubinger/aileron-connector-linear"
op = "update_integration_slack_scopes"
idempotent = false

[approval]
required = true

[[inputs]]
name = "code"
type = "String"
description = "The Slack OAuth code."
required = true

[[inputs]]
name = "integrationId"
type = "String"
description = "The ID of the existing Slack integration"
required = true

[[inputs]]
name = "redirectUri"
type = "String"
description = "The Slack OAuth redirect URI."
required = true
+++
