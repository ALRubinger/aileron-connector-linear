+++
name = "customer-need"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/customer-need@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["customer_need"]

[match]
intent = "Retrieves a single customer need by ID or hash."

[[execute]]
id = "customer"
connector = "github://ALRubinger/aileron-connector-linear"
op = "customer_need"
idempotent = true

[[inputs]]
name = "hash"
type = "String"
description = "The hash prefix of the customer need to retrieve (minimum 8 characters). Cannot be used together with id."
required = false

[[inputs]]
name = "id"
type = "String"
description = "The UUID of the customer need to retrieve. Cannot be used together with hash."
required = false
+++
