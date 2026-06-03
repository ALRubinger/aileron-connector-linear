+++
name = "issue-import-create-jira"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/issue-import-create-jira@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["issue_import_create_jira"]

[match]
intent = "Kicks off a Jira import job."

[[execute]]
id = "issue"
connector = "github://ALRubinger/aileron-connector-linear"
op = "issue_import_create_jira"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "ID of issue import. If not provided it will be generated."
required = false

[[inputs]]
name = "includeClosedIssues"
type = "Boolean"
description = "Whether or not we should collect the data for closed issues."
required = false

[[inputs]]
name = "instantProcess"
type = "Boolean"
description = "Whether to instantly process the import with the default configuration mapping."
required = false

[[inputs]]
name = "jiraEmail"
type = "String"
description = "Jira user account email."
required = true

[[inputs]]
name = "jiraHostname"
type = "String"
description = "Jira installation or cloud hostname."
required = true

[[inputs]]
name = "jiraProject"
type = "String"
description = "Jira project key from which we will import data."
required = true

[[inputs]]
name = "jiraToken"
type = "String"
description = "Jira personal access token to access Jira REST API."
required = true

[[inputs]]
name = "jql"
type = "String"
description = "A custom JQL query to filter issues being imported"
required = false

[[inputs]]
name = "teamId"
type = "String"
description = "ID of the team into which to import data. Empty to create new team."
required = false

[[inputs]]
name = "teamName"
type = "String"
description = "Name of new team. When teamId is not set."
required = false
+++
