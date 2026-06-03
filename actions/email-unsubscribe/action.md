+++
name = "email-unsubscribe"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/email-unsubscribe@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["email_unsubscribe"]

[match]
intent = "Unsubscribes the user from one type of email."

[[execute]]
id = "email"
connector = "github://ALRubinger/aileron-connector-linear"
op = "email_unsubscribe"
idempotent = false

[approval]
required = true

[[inputs]]
name = "token"
type = "String"
description = "The user's email validation token."
required = true

[[inputs]]
name = "type"
type = "String"
description = "Email type to unsubscribe from."
required = true

[[inputs]]
name = "userId"
type = "String"
description = "The identifier of the user."
required = true
+++
