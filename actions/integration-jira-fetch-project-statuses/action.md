+++
name = "integration-jira-fetch-project-statuses"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/integration-jira-fetch-project-statuses@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["integration_jira_fetch_project_statuses"]

[match]
intent = "[INTERNAL] Fetches Jira project statuses and stores them in integration settings."

[[execute]]
id = "integration"
connector = "github://ALRubinger/aileron-connector-linear"
op = "integration_jira_fetch_project_statuses"
idempotent = false

[approval]
required = true

[[inputs]]
name = "integrationId"
type = "String"
description = "The id of the Jira integration."
required = true

[[inputs]]
name = "projectId"
type = "String"
description = "The Jira project ID to fetch statuses for."
required = true
+++
