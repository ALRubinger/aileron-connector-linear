+++
name = "jira-integration-connect"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/jira-integration-connect@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["jira_integration_connect"]

[match]
intent = "[INTERNAL] Connects the workspace with a Jira Personal Access Token."

[[execute]]
id = "jira"
connector = "github://ALRubinger/aileron-connector-linear"
op = "jira_integration_connect"
idempotent = false

[approval]
required = true

[[inputs]]
name = "accessToken"
type = "String"
description = "The Jira personal access token."
required = true

[[inputs]]
name = "email"
type = "String"
description = "The Jira user's email address. A username is also accepted on Jira Server / DC."
required = true

[[inputs]]
name = "hostname"
type = "String"
description = "The Jira installation hostname."
required = true

[[inputs]]
name = "manualSetup"
type = "Boolean"
description = "Whether this integration will be setup using the manual webhook flow."
required = false
+++
