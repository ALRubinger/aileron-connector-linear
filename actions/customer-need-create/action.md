+++
name = "customer-need-create"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/customer-need-create@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["customer_need_create"]

[match]
intent = "Creates a new customer need."

[[execute]]
id = "customer"
connector = "github://ALRubinger/aileron-connector-linear"
op = "customer_need_create"
idempotent = false

[approval]
required = true

[[inputs]]
name = "attachmentId"
type = "String"
description = "The UUID of an existing attachment to associate with this need as its source."
required = false

[[inputs]]
name = "attachmentUrl"
type = "String"
description = "A URL to create an attachment from and associate with this customer need as its source."
required = false

[[inputs]]
name = "body"
type = "String"
description = "The body content of the need in Markdown format. Cannot be used together with bodyData."
required = false

[[inputs]]
name = "bodyData"
type = "JSON"
description = "[Internal] The body content of the need as a Prosemirror document JSON string. Cannot be used together with body."
required = false

[[inputs]]
name = "commentId"
type = "String"
description = "The UUID of an existing comment to associate with this need for additional context."
required = false

[[inputs]]
name = "createAsUser"
type = "String"
description = "Create the need attributed to an external user with the provided name. This option is only available to OAuth applications creating needs in `actor=app` mode."
required = false

[[inputs]]
name = "createdAt"
type = "DateTime"
description = "The time at which the customer need was created (e.g. if importing from another system). Must be a time in the past. If none is provided, the backend will generate the time as now."
required = false

[[inputs]]
name = "customerExternalId"
type = "String"
description = "The external system ID of the customer this need belongs to. Cannot be used together with customerId."
required = false

[[inputs]]
name = "customerId"
type = "String"
description = "The UUID of the customer this need belongs to. Cannot be used together with customerExternalId."
required = false

[[inputs]]
name = "displayIconUrl"
type = "String"
description = "Avatar URL for the external user specified in `createAsUser`. Can only be used in conjunction with `createAsUser`. This option is only available to OAuth applications creating needs in `actor=app` mode."
required = false

[[inputs]]
name = "id"
type = "String"
description = "The identifier in UUID v4 format. If none is provided, the backend will generate one."
required = false

[[inputs]]
name = "issueId"
type = "String"
description = "The issue to link this need to. Accepts a UUID or issue identifier (e.g., 'LIN-123'). Either issueId or projectId must be provided."
required = false

[[inputs]]
name = "priority"
type = "Float"
description = "Whether the customer need is important or not. 0 = Not important, 1 = Important."
required = false

[[inputs]]
name = "projectId"
type = "String"
description = "[INTERNAL] The project to link this need to. Either issueId or projectId must be provided."
required = false
+++
