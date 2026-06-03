+++
name = "customer"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/customer@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["customer"]

[match]
intent = "Retrieves a single customer by ID or slug."

[[execute]]
id = "customer"
connector = "github://ALRubinger/aileron-connector-linear"
op = "customer"
idempotent = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier or slug of the customer to retrieve."
required = true
+++
