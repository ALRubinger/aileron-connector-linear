+++
name = "customer-need-unarchive"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/customer-need-unarchive@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["customer_need_unarchive"]

[match]
intent = "Unarchives a customer need."

[[execute]]
id = "customer"
connector = "github://ALRubinger/aileron-connector-linear"
op = "customer_need_unarchive"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the customer need to unarchive."
required = true
+++
