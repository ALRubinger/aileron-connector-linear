+++
name = "user-external-user-disconnect"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/user-external-user-disconnect@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["user_external_user_disconnect"]

[match]
intent = "Disconnects the external user from this Linear account."

[[execute]]
id = "user"
connector = "github://ALRubinger/aileron-connector-linear"
op = "user_external_user_disconnect"
idempotent = false

[approval]
required = true

[[inputs]]
name = "service"
type = "String"
description = "The external service to disconnect."
required = true
+++
