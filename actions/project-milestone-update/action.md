+++
name = "project-milestone-update"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/project-milestone-update@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["project_milestone_update"]

[match]
intent = "Updates a project milestone."

[[execute]]
id = "project"
connector = "github://ALRubinger/aileron-connector-linear"
op = "project_milestone_update"
idempotent = false

[approval]
required = true

[[inputs]]
name = "description"
type = "String"
description = "The description of the project milestone in markdown format."
required = false

[[inputs]]
name = "descriptionData"
type = "JSONObject"
description = "[Internal] The description of the project milestone as a Prosemirror document."
required = false

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the project milestone to update. Also the identifier from the URL is accepted."
required = true

[[inputs]]
name = "name"
type = "String"
description = "The name of the project milestone."
required = false

[[inputs]]
name = "projectId"
type = "String"
description = "Related project for the project milestone."
required = false

[[inputs]]
name = "sortOrder"
type = "Float"
description = "The sort order for the project milestone within a project."
required = false

[[inputs]]
name = "targetDate"
type = "TimelessDate"
description = "The planned target date of the project milestone."
required = false
+++
