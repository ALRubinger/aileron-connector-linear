+++
name = "user-updated"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/user-updated@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["user_updated"]

[match]
intent = "Triggered when an user is updated"

[[execute]]
id = "user"
connector = "github://ALRubinger/aileron-connector-linear"
op = "user_updated"
idempotent = true
+++
