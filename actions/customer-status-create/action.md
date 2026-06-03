+++
name = "customer-status-create"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/customer-status-create@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["customer_status_create"]

[match]
intent = "Creates a new customer status."

[[execute]]
id = "customer"
connector = "github://ALRubinger/aileron-connector-linear"
op = "customer_status_create"
idempotent = false

[approval]
required = true

[[inputs]]
name = "color"
type = "String"
description = "The color of the status indicator in the UI, as a HEX string (e.g., '#ff0000')."
required = true

[[inputs]]
name = "description"
type = "String"
description = "An optional description explaining what this status represents."
required = false

[[inputs]]
name = "displayName"
type = "String"
description = "The user-facing display name of the status. At least one of name or displayName must be provided."
required = false

[[inputs]]
name = "id"
type = "String"
description = "The identifier in UUID v4 format. If none is provided, the backend will generate one."
required = false

[[inputs]]
name = "name"
type = "String"
description = "The internal name of the status. At least one of name or displayName must be provided."
required = false

[[inputs]]
name = "position"
type = "Float"
description = "The sort position of the status in the workspace's customer lifecycle flow. If omitted or colliding, a position is automatically assigned at the end."
required = false
+++
