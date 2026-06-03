+++
name = "initiative-updated"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/initiative-updated@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["initiative_updated"]

[match]
intent = "Triggered when an initiative is updated"

[[execute]]
id = "initiative"
connector = "github://ALRubinger/aileron-connector-linear"
op = "initiative_updated"
idempotent = true
+++
