+++
name = "integration-slack-or-asks-update-slack-team-name"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/integration-slack-or-asks-update-slack-team-name@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["integration_slack_or_asks_update_slack_team_name"]

[match]
intent = "Updates the Slack team's name in Linear for an existing Slack or Asks integration."

[[execute]]
id = "integration"
connector = "github://ALRubinger/aileron-connector-linear"
op = "integration_slack_or_asks_update_slack_team_name"
idempotent = false

[approval]
required = true

[[inputs]]
name = "integrationId"
type = "String"
description = "The integration ID."
required = true
+++
