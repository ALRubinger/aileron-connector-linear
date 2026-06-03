+++
name = "logout-session"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/logout-session@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["logout_session"]

[match]
intent = "Logout an individual session with its ID."

[[execute]]
id = "logout"
connector = "github://ALRubinger/aileron-connector-linear"
op = "logout_session"
idempotent = false

[approval]
required = true

[[inputs]]
name = "sessionId"
type = "String"
description = "ID of the session to logout."
required = true
+++
