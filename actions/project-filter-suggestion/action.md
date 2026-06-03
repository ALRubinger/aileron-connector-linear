+++
name = "project-filter-suggestion"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/project-filter-suggestion@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["project_filter_suggestion"]

[match]
intent = "Suggests filters for a project view based on a text prompt."

[[execute]]
id = "project"
connector = "github://ALRubinger/aileron-connector-linear"
op = "project_filter_suggestion"
idempotent = true

[[inputs]]
name = "prompt"
type = "String"
description = ""
required = true

[[inputs]]
name = "teamId"
type = "String"
description = "[Internal] The ID of the team if filtering a team view."
required = false
+++
