+++
name = "authentication-sessions"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/authentication-sessions@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["authentication_sessions"]

[match]
intent = "User's active sessions."

[[execute]]
id = "authentication"
connector = "github://ALRubinger/aileron-connector-linear"
op = "authentication_sessions"
idempotent = true
+++
