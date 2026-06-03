+++
name = "logout"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/logout@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["logout"]

[match]
intent = "Logout the client."

[[execute]]
id = "logout"
connector = "github://ALRubinger/aileron-connector-linear"
op = "logout"
idempotent = false

[approval]
required = true

[[inputs]]
name = "reason"
type = "String"
description = "The reason for logging out."
required = false
+++
