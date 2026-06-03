+++
name = "favorite-created"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/favorite-created@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["favorite_created"]

[match]
intent = "Triggered when a favorite is created"

[[execute]]
id = "favorite"
connector = "github://ALRubinger/aileron-connector-linear"
op = "favorite_created"
idempotent = true
+++
