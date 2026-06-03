+++
name = "project-label-update"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/project-label-update@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["project_label_update"]

[match]
intent = "Updates a project label."

[[execute]]
id = "project"
connector = "github://ALRubinger/aileron-connector-linear"
op = "project_label_update"
idempotent = false

[approval]
required = true

[[inputs]]
name = "color"
type = "String"
description = "The color of the label."
required = false

[[inputs]]
name = "description"
type = "String"
description = "The description of the label."
required = false

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the label to update."
required = true

[[inputs]]
name = "isGroup"
type = "Boolean"
description = "Whether the label is a group."
required = false

[[inputs]]
name = "name"
type = "String"
description = "The name of the label."
required = false

[[inputs]]
name = "parentId"
type = "String"
description = "The identifier of the parent label."
required = false

[[inputs]]
name = "retiredAt"
type = "DateTime"
description = "The time at which the label was retired. Set to null to restore a retired label."
required = false
+++
