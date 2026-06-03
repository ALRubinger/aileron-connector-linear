+++
name = "user-created"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/user-created@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["user_created"]

[match]
intent = "Triggered when an user is created"

[[execute]]
id = "user"
connector = "github://ALRubinger/aileron-connector-linear"
op = "user_created"
idempotent = true
+++
