+++
name = "customer-status-delete"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/customer-status-delete@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["customer_status_delete"]

[match]
intent = "Deletes a customer status. Cannot delete the last remaining status in a workspace, and the status must not be in use by any customers."

[[execute]]
id = "customer"
connector = "github://ALRubinger/aileron-connector-linear"
op = "customer_status_delete"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the customer status to delete."
required = true
+++
