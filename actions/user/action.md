+++
name = "user"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/user@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["user"]

[match]
intent = "Fetches a specific user by their ID."

[[execute]]
id = "user"
connector = "github://ALRubinger/aileron-connector-linear"
op = "user"
idempotent = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the user to retrieve. To retrieve the authenticated user, use `viewer` query."
required = true
+++
