+++
name = "customer-need-delete"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/customer-need-delete@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["customer_need_delete"]

[match]
intent = "Deletes a customer need."

[[execute]]
id = "customer"
connector = "github://ALRubinger/aileron-connector-linear"
op = "customer_need_delete"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the customer need to delete."
required = true

[[inputs]]
name = "keepAttachment"
type = "Boolean"
description = "Whether to keep the attachment associated with the customer need."
required = false
+++
