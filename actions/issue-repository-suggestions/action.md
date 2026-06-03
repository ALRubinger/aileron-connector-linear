+++
name = "issue-repository-suggestions"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/issue-repository-suggestions@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["issue_repository_suggestions"]

[match]
intent = "Returns code repositories that are most likely to be relevant for implementing an issue."

[[execute]]
id = "issue"
connector = "github://ALRubinger/aileron-connector-linear"
op = "issue_repository_suggestions"
idempotent = true

[[inputs]]
name = "agentSessionId"
type = "String"
description = "Optional AgentSession ID associated with the issue for which the suggestions are being generated."
required = false

[[inputs]]
name = "hostname"
type = "String"
description = "Hostname of the Git service (e.g., 'github.com', 'github.company.com')."
required = true

[[inputs]]
name = "issueId"
type = "String"
description = "The ID of the issue to get repository suggestions for."
required = true

[[inputs]]
name = "repositoryFullName"
type = "String"
description = "The full name of the repository in owner/name format (e.g., 'acme/backend')."
required = true
+++
