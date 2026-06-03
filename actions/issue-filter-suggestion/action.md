+++
name = "issue-filter-suggestion"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/issue-filter-suggestion@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["issue_filter_suggestion"]

[match]
intent = "Suggests filters for an issue view based on a text prompt."

[[execute]]
id = "issue"
connector = "github://ALRubinger/aileron-connector-linear"
op = "issue_filter_suggestion"
idempotent = true

[[inputs]]
name = "projectId"
type = "String"
description = "The ID of the project if filtering a project view."
required = false

[[inputs]]
name = "prompt"
type = "String"
description = "The text prompt to generate a filter suggestion from."
required = true

[[inputs]]
name = "teamId"
type = "String"
description = "[Internal] The ID of the team if filtering a team view."
required = false
+++
