+++
name = "issue-label-create"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/issue-label-create@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["issue_label_create"]

[match]
intent = "Creates a new label."

[[execute]]
id = "issue"
connector = "github://ALRubinger/aileron-connector-linear"
op = "issue_label_create"
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
description = "The identifier in UUID v4 format. If none is provided, the backend will generate one."
required = false

[[inputs]]
name = "isGroup"
type = "Boolean"
description = "Whether the label is a group."
required = false

[[inputs]]
name = "name"
type = "String"
description = "The name of the label."
required = true

[[inputs]]
name = "parentId"
type = "String"
description = "The identifier of the parent label."
required = false

[[inputs]]
name = "replaceTeamLabels"
type = "Boolean"
description = "Whether to replace all team-specific labels with the same name with this newly created workspace label (default: false)."
required = false

[[inputs]]
name = "retiredAt"
type = "DateTime"
description = "The time at which the label was retired. Set to null to restore a retired label."
required = false

[[inputs]]
name = "teamId"
type = "String"
description = "The team associated with the label. If not given, the label will be associated with the entire workspace."
required = false
+++
