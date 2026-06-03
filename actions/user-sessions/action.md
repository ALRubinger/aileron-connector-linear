+++
name = "user-sessions"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/user-sessions@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["user_sessions"]

[match]
intent = "Lists all active authentication sessions for a user. Can only be called by a workspace admin or owner."

[[execute]]
id = "user"
connector = "github://ALRubinger/aileron-connector-linear"
op = "user_sessions"
idempotent = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the user to list sessions of."
required = true
+++
