+++
name = "logout-all-sessions"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/logout-all-sessions@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["logout_all_sessions"]

[match]
intent = "Logout all of user's sessions including the active one."

[[execute]]
id = "logout"
connector = "github://ALRubinger/aileron-connector-linear"
op = "logout_all_sessions"
idempotent = false

[approval]
required = true

[[inputs]]
name = "reason"
type = "String"
description = "The reason for logging out."
required = false
+++
