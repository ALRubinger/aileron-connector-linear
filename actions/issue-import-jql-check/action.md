+++
name = "issue-import-jql-check"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/issue-import-jql-check@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["issue_import_jql_check"]

[match]
intent = "Checks whether a custom JQL query is valid and can be used to filter issues of a Jira import."

[[execute]]
id = "issue"
connector = "github://ALRubinger/aileron-connector-linear"
op = "issue_import_jql_check"
idempotent = true

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
description = "Jira project key to use as the base filter of the query."
required = true

[[inputs]]
name = "jiraToken"
type = "String"
description = "Jira personal access token to access Jira REST API."
required = true

[[inputs]]
name = "jql"
type = "String"
description = "The JQL query to validate."
required = true
+++
