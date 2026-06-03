+++
name = "customer-merge"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/customer-merge@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["customer_merge"]

[match]
intent = "Merges two customers by transferring all needs from the source customer to the target customer. The source customer is archived after the merge. Domains, external IDs, and metadata are combined on the target customer."

[[execute]]
id = "customer"
connector = "github://ALRubinger/aileron-connector-linear"
op = "customer_merge"
idempotent = false

[approval]
required = true

[[inputs]]
name = "sourceCustomerId"
type = "String"
description = "The ID of the source customer to merge. All needs from this customer will be transferred to the target, and this customer will be archived."
required = true

[[inputs]]
name = "targetCustomerId"
type = "String"
description = "The ID of the target customer to merge into. This customer will receive all needs from the source customer and retain its own."
required = true
+++
