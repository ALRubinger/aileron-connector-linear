+++
name = "customer-tier-update"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/customer-tier-update@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["customer_tier_update"]

[match]
intent = "Updates a customer tier."

[[execute]]
id = "customer"
connector = "github://ALRubinger/aileron-connector-linear"
op = "customer_tier_update"
idempotent = false

[approval]
required = true

[[inputs]]
name = "color"
type = "String"
description = "The updated color of the tier indicator in the UI, as a HEX string."
required = false

[[inputs]]
name = "description"
type = "String"
description = "The updated description of the tier."
required = false

[[inputs]]
name = "displayName"
type = "String"
description = "The updated user-facing display name of the tier."
required = false

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the customer tier to update."
required = true

[[inputs]]
name = "name"
type = "String"
description = "The updated internal name of the tier."
required = false

[[inputs]]
name = "position"
type = "Float"
description = "The updated sort position of the tier in the workspace's customer tier ordering."
required = false
+++
