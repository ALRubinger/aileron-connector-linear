+++
name = "integration-jira-update"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/integration-jira-update@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["integration_jira_update"]

[match]
intent = "[INTERNAL] Updates a Jira Integration."

[[execute]]
id = "integration"
connector = "github://ALRubinger/aileron-connector-linear"
op = "integration_jira_update"
idempotent = false

[approval]
required = true

[[inputs]]
name = "accessToken"
type = "String"
description = "The Jira personal access token."
required = false

[[inputs]]
name = "deleteWebhook"
type = "Boolean"
description = "Whether to delete the current manual webhook configuration."
required = false

[[inputs]]
name = "email"
type = "String"
description = "The Jira user email address associated with the personal access token."
required = false

[[inputs]]
name = "id"
type = "String"
description = "The id of the integration to update."
required = true

[[inputs]]
name = "noSecret"
type = "Boolean"
description = "Whether the Jira instance does not support webhook secrets."
required = false

[[inputs]]
name = "updateMetadata"
type = "Boolean"
description = "Whether to refresh Jira metadata for the integration."
required = false

[[inputs]]
name = "updateProjects"
type = "Boolean"
description = "Whether to refresh Jira Projects for the integration."
required = false

[[inputs]]
name = "webhookSecret"
type = "String"
description = "Webhook secret for a new manual configuration."
required = false
+++
