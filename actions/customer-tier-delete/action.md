+++
name = "customer-tier-delete"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/customer-tier-delete@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["customer_tier_delete"]

[match]
intent = "Deletes a customer tier. The tier must not be in use by any customers."

[[execute]]
id = "customer"
connector = "github://ALRubinger/aileron-connector-linear"
op = "customer_tier_delete"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the customer tier to delete."
required = true
+++
