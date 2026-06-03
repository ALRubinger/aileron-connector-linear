+++
name = "customer-unsync"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/customer-unsync@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["customer_unsync"]

[match]
intent = "Unsyncs a managed customer from its current data source integration. External IDs mapping to the external source will be cleared, and the customer will no longer be updated by the integration."

[[execute]]
id = "customer"
connector = "github://ALRubinger/aileron-connector-linear"
op = "customer_unsync"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the customer to unsync."
required = true
+++
