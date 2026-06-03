+++
name = "user-revoke-session"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/user-revoke-session@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["user_revoke_session"]

[match]
intent = "Revokes a specific authentication session for a user. The admin cannot revoke their own sessions with this mutation. Can only be called by a workspace admin or owner."

[[execute]]
id = "user"
connector = "github://ALRubinger/aileron-connector-linear"
op = "user_revoke_session"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the user whose session to revoke."
required = true

[[inputs]]
name = "sessionId"
type = "String"
description = "The identifier of the session to revoke."
required = true
+++
