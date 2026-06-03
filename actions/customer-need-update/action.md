+++
name = "customer-need-update"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/customer-need-update@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["customer_need_update"]

[match]
intent = "Updates an existing customer need. Supports moving the need to a different issue or project, changing priority, updating the body content, and managing the attached source URL."

[[execute]]
id = "customer"
connector = "github://ALRubinger/aileron-connector-linear"
op = "customer_need_update"
idempotent = false

[approval]
required = true

[[inputs]]
name = "applyPriorityToRelatedNeeds"
type = "Boolean"
description = "When true and priority is also set, applies the same priority update to all other needs from the same customer on the same issue or project."
required = false

[[inputs]]
name = "attachmentUrl"
type = "String"
description = "A URL to create a new attachment from and set as the source for this customer need. Replaces any existing manually-added attachment."
required = false

[[inputs]]
name = "body"
type = "String"
description = "The updated body content of the need in Markdown format. Set to null to clear the body. Cannot be used together with bodyData."
required = false

[[inputs]]
name = "bodyData"
type = "JSON"
description = "[Internal] The updated body content of the need as a Prosemirror document JSON string. Set to null to clear the body. Cannot be used together with body."
required = false

[[inputs]]
name = "clearAttachment"
type = "Boolean"
description = "Whether to remove any existing attachment associated with the customer need."
required = false

[[inputs]]
name = "customerExternalId"
type = "String"
description = "The external system ID of the customer to reassign this need to. Cannot be used together with customerId."
required = false

[[inputs]]
name = "customerId"
type = "String"
description = "The UUID of the customer to reassign this need to. Cannot be used together with customerExternalId."
required = false

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the customer need to update."
required = true

[[inputs]]
name = "id"
type = "String"
description = "An optional identifier in UUID v4 format. If provided, will be set as the customer need's ID."
required = false

[[inputs]]
name = "issueId"
type = "String"
description = "The issue to move this need to. Accepts a UUID or issue identifier (e.g., 'LIN-123'). The need's attachment will also be moved to the target issue."
required = false

[[inputs]]
name = "priority"
type = "Float"
description = "Whether the customer need is important or not. 0 = Not important, 1 = Important."
required = false

[[inputs]]
name = "projectId"
type = "String"
description = "[INTERNAL] The project to move this need to."
required = false
+++
