+++
name = "template-update"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/template-update@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["template_update"]

[match]
intent = "Updates an existing template."

[[execute]]
id = "template"
connector = "github://ALRubinger/aileron-connector-linear"
op = "template_update"
idempotent = false

[approval]
required = true

[[inputs]]
name = "color"
type = "String"
description = "The color of the template icon."
required = false

[[inputs]]
name = "description"
type = "String"
description = "The template description."
required = false

[[inputs]]
name = "icon"
type = "String"
description = "The icon of the template."
required = false

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the template."
required = true

[[inputs]]
name = "name"
type = "String"
description = "The template name."
required = false

[[inputs]]
name = "sortOrder"
type = "Float"
description = "The position of the template in the templates list."
required = false

[[inputs]]
name = "teamId"
type = "String"
description = "The identifier or key of the team associated with the template. If set to null, the template will be shared across all teams."
required = false

[[inputs]]
name = "templateData"
type = "JSON"
description = "The template data as JSON-encoded attributes of the target entity type, such as pre-filled issue fields, project configuration, or document content."
required = false
+++
