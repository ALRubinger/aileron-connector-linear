+++
name = "customer-status"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/customer-status@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["customer_status"]

[match]
intent = "One specific customer status."

[[execute]]
id = "customer"
connector = "github://ALRubinger/aileron-connector-linear"
op = "customer_status"
idempotent = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the customer status to retrieve."
required = true
+++
