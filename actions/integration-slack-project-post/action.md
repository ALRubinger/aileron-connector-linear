+++
name = "integration-slack-project-post"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/integration-slack-project-post@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["integration_slack_project_post"]

[match]
intent = "Slack integration for project notifications."

[[execute]]
id = "integration"
connector = "github://ALRubinger/aileron-connector-linear"
op = "integration_slack_project_post"
idempotent = false

[approval]
required = true

[[inputs]]
name = "code"
type = "String"
description = "The Slack OAuth code."
required = true

[[inputs]]
name = "projectId"
type = "String"
description = "Integration's associated project."
required = true

[[inputs]]
name = "redirectUri"
type = "String"
description = "The Slack OAuth redirect URI."
required = true

[[inputs]]
name = "service"
type = "String"
description = "The service to enable once connected, either 'notifications' or 'updates'."
required = true
+++
