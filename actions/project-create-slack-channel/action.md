+++
name = "project-create-slack-channel"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/project-create-slack-channel@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["project_create_slack_channel"]

[match]
intent = "[Internal] Creates a Slack channel for an existing project."

[[execute]]
id = "project"
connector = "github://ALRubinger/aileron-connector-linear"
op = "project_create_slack_channel"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the project."
required = true

[[inputs]]
name = "integrationId"
type = "String"
description = "The identifier of the Slack integration to use. When not provided, uses the organization's configured integration."
required = false

[[inputs]]
name = "slackChannelName"
type = "String"
description = "The full name for the Slack channel to create (including prefix). When provided, a Slack channel will be created and connected to the project."
required = true
+++
