+++
name = "integration-jira-personal"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/integration-jira-personal@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["integration_jira_personal"]

[match]
intent = "Connect your Jira account to Linear."

[[execute]]
id = "integration"
connector = "github://ALRubinger/aileron-connector-linear"
op = "integration_jira_personal"
idempotent = false

[approval]
required = true

[[inputs]]
name = "accessToken"
type = "String"
description = "The Jira personal access token, when connecting using a PAT."
required = false

[[inputs]]
name = "code"
type = "String"
description = "The Jira OAuth code, when connecting using OAuth."
required = false
+++
