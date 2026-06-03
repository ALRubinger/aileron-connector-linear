+++
name = "integration-slack-custom-view-notifications"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/integration-slack-custom-view-notifications@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["integration_slack_custom_view_notifications"]

[match]
intent = "Slack integration for custom view notifications."

[[execute]]
id = "integration"
connector = "github://ALRubinger/aileron-connector-linear"
op = "integration_slack_custom_view_notifications"
idempotent = false

[approval]
required = true

[[inputs]]
name = "code"
type = "String"
description = "The Slack OAuth code."
required = true

[[inputs]]
name = "customViewId"
type = "String"
description = "Integration's associated custom view."
required = true

[[inputs]]
name = "redirectUri"
type = "String"
description = "The Slack OAuth redirect URI."
required = true
+++
