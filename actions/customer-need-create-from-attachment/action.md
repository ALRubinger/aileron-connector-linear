+++
name = "customer-need-create-from-attachment"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/customer-need-create-from-attachment@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["customer_need_create_from_attachment"]

[match]
intent = "Creates a new customer need from an existing issue attachment. If the attachment already has an archived need, it will be unarchived instead of creating a duplicate."

[[execute]]
id = "customer"
connector = "github://ALRubinger/aileron-connector-linear"
op = "customer_need_create_from_attachment"
idempotent = false

[approval]
required = true

[[inputs]]
name = "attachmentId"
type = "String"
description = "The UUID of the existing issue attachment to create a customer need from."
required = true
+++
