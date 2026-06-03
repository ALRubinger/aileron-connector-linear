+++
name = "customer-tier-create"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/customer-tier-create@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["customer_tier_create"]

[match]
intent = "Creates a new customer tier."

[[execute]]
id = "customer"
connector = "github://ALRubinger/aileron-connector-linear"
op = "customer_tier_create"
idempotent = false

[approval]
required = true

[[inputs]]
name = "color"
type = "String"
description = "The color of the tier indicator in the UI, as a HEX string (e.g., '#ff0000')."
required = true

[[inputs]]
name = "description"
type = "String"
description = "An optional description explaining what this tier represents."
required = false

[[inputs]]
name = "displayName"
type = "String"
description = "The user-facing display name of the tier. At least one of name or displayName must be provided."
required = false

[[inputs]]
name = "id"
type = "String"
description = "The identifier in UUID v4 format. If none is provided, the backend will generate one."
required = false

[[inputs]]
name = "name"
type = "String"
description = "The internal name of the tier. Must be unique within the workspace. At least one of name or displayName must be provided."
required = false

[[inputs]]
name = "position"
type = "Float"
description = "The sort position of the tier in the workspace's customer tier ordering. If omitted or colliding, a position is automatically assigned at the end."
required = false
+++
