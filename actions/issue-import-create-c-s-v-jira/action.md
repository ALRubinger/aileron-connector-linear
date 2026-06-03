+++
name = "issue-import-create-c-s-v-jira"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/issue-import-create-c-s-v-jira@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["issue_import_create_c_s_v_jira"]

[match]
intent = "Kicks off a Jira import job from a CSV."

[[execute]]
id = "issue"
connector = "github://ALRubinger/aileron-connector-linear"
op = "issue_import_create_c_s_v_jira"
idempotent = false

[approval]
required = true

[[inputs]]
name = "csvUrl"
type = "String"
description = "URL for the CSV."
required = true

[[inputs]]
name = "jiraEmail"
type = "String"
description = "Jira user account email."
required = false

[[inputs]]
name = "jiraHostname"
type = "String"
description = "Jira installation or cloud hostname."
required = false

[[inputs]]
name = "jiraToken"
type = "String"
description = "Jira personal access token to access Jira REST API."
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
