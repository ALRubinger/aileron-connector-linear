+++
name = "email-intake-address"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/email-intake-address@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["email_intake_address"]

[match]
intent = "One specific email intake address."

[[execute]]
id = "email"
connector = "github://ALRubinger/aileron-connector-linear"
op = "email_intake_address"
idempotent = true

[[inputs]]
name = "id"
type = "String"
description = ""
required = true
+++
