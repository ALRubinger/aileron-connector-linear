+++
name = "integration-slack-workflow-access-update"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/integration-slack-workflow-access-update@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["integration_slack_workflow_access_update"]

[match]
intent = "[Internal] Enables Linear Agent Slack workflow access for a Slack or Slack Asks integration."

[[execute]]
id = "integration"
connector = "github://ALRubinger/aileron-connector-linear"
op = "integration_slack_workflow_access_update"
idempotent = false

[approval]
required = true

[[inputs]]
name = "enabled"
type = "Boolean"
description = "Whether to enable workflow access."
required = true

[[inputs]]
name = "integrationId"
type = "String"
description = "The ID of the integration to toggle workflow access for."
required = true
+++
