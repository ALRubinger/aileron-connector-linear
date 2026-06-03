+++
name = "project-relation-create"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/project-relation-create@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["project_relation_create"]

[match]
intent = "Creates a new project relation."

[[execute]]
id = "project"
connector = "github://ALRubinger/aileron-connector-linear"
op = "project_relation_create"
idempotent = false

[approval]
required = true

[[inputs]]
name = "anchorType"
type = "String"
description = "The type of the anchor for the project."
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier in UUID v4 format. If none is provided, the backend will generate one."
required = false

[[inputs]]
name = "projectId"
type = "String"
description = "The identifier of the project that is related to another project."
required = true

[[inputs]]
name = "projectMilestoneId"
type = "String"
description = "The identifier of the project milestone."
required = false

[[inputs]]
name = "relatedAnchorType"
type = "String"
description = "The type of the anchor for the related project."
required = true

[[inputs]]
name = "relatedProjectId"
type = "String"
description = "The identifier of the related project."
required = true

[[inputs]]
name = "relatedProjectMilestoneId"
type = "String"
description = "The identifier of the related project milestone."
required = false

[[inputs]]
name = "type"
type = "String"
description = "The type of relation of the project to the related project."
required = true
+++
