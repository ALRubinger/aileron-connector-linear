+++
name = "integration-slack-org-project-updates-post"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/integration-slack-org-project-updates-post@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["integration_slack_org_project_updates_post"]

[match]
intent = "Slack integration for workspace-level project update notifications."

[[execute]]
id = "integration"
connector = "github://ALRubinger/aileron-connector-linear"
op = "integration_slack_org_project_updates_post"
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
