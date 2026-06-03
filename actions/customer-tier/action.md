+++
name = "customer-tier"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/customer-tier@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["customer_tier"]

[match]
intent = "One specific customer tier."

[[execute]]
id = "customer"
connector = "github://ALRubinger/aileron-connector-linear"
op = "customer_tier"
idempotent = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the customer tier to retrieve."
required = true
+++
