+++
name = "user-revoke-all-sessions"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/user-revoke-all-sessions@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["user_revoke_all_sessions"]

[match]
intent = "Revokes all active sessions for a user, forcing them to re-authenticate. Can only be called by a workspace admin or owner."

[[execute]]
id = "user"
connector = "github://ALRubinger/aileron-connector-linear"
op = "user_revoke_all_sessions"
idempotent = false

[approval]
required = true

[[inputs]]
name = "id"
type = "String"
description = "The identifier of the user to logout all sessions of."
required = true
+++
