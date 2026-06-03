+++
name = "initiative-create"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/initiative-create@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["initiative_create"]

[match]
intent = "Creates a new initiative."

[[execute]]
id = "initiative"
connector = "github://ALRubinger/aileron-connector-linear"
op = "initiative_create"
idempotent = false

[approval]
required = true

[[inputs]]
name = "color"
type = "String"
description = "The initiative's color."
required = false

[[inputs]]
name = "content"
type = "String"
description = "The initiative's content in markdown format."
required = false

[[inputs]]
name = "description"
type = "String"
description = "The description of the initiative."
required = false

[[inputs]]
name = "icon"
type = "String"
description = "The initiative's icon."
required = false

[[inputs]]
name = "id"
type = "String"
description = "The identifier in UUID v4 format. If none is provided, the backend will generate one."
required = false

[[inputs]]
name = "labelIds"
type = "String"
description = "[Internal] The identifiers of the initiative labels associated with this initiative."
required = false

[[inputs]]
name = "name"
type = "String"
description = "The name of the initiative."
required = true

[[inputs]]
name = "ownerId"
type = "String"
description = "The owner of the initiative."
required = false

[[inputs]]
name = "sortOrder"
type = "Float"
description = "The sort order of the initiative within the workspace."
required = false

[[inputs]]
name = "status"
type = "InitiativeStatus"
description = "The initiative's status."
required = false

[[inputs]]
name = "targetDate"
type = "TimelessDate"
description = "The estimated completion date of the initiative."
required = false

[[inputs]]
name = "targetDateResolution"
type = "DateResolutionType"
description = "The resolution of the initiative's estimated completion date."
required = false
+++
