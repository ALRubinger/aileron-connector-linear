+++
name = "entity-external-link-create"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/entity-external-link-create@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["entity_external_link_create"]

[match]
intent = "Creates a new external link on an initiative, project, team, release, or cycle."

[[execute]]
id = "entity"
connector = "github://ALRubinger/aileron-connector-linear"
op = "entity_external_link_create"
idempotent = false

[approval]
required = true

[[inputs]]
name = "cycleId"
type = "String"
description = "[Internal] The cycle associated with the link."
required = false

[[inputs]]
name = "id"
type = "String"
description = "The identifier in UUID v4 format. If none is provided, the backend will generate one."
required = false

[[inputs]]
name = "initiativeId"
type = "String"
description = "The initiative associated with the link."
required = false

[[inputs]]
name = "label"
type = "String"
description = "The label for the link."
required = true

[[inputs]]
name = "projectId"
type = "String"
description = "The project associated with the link."
required = false

[[inputs]]
name = "releaseId"
type = "String"
description = "The release associated with the link."
required = false

[[inputs]]
name = "resourceFolderId"
type = "String"
description = "[Internal] The resource folder containing the link."
required = false

[[inputs]]
name = "sortOrder"
type = "Float"
description = "The order of the item in the entities resources list."
required = false

[[inputs]]
name = "teamId"
type = "String"
description = "[Internal] The team associated with the link."
required = false

[[inputs]]
name = "url"
type = "String"
description = "The URL of the link."
required = true
+++
