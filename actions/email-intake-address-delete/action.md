+++
name = "email-intake-address-delete"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/email-intake-address-delete@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["email_intake_address_delete"]

[match]
intent = "Deletes an email intake address object."

[[execute]]
id = "email"
connector = "github://ALRubinger/aileron-connector-linear"
op = "email_intake_address_delete"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the email intake address to delete."
required = true
+++
