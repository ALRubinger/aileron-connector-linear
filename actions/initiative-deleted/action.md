+++
name = "initiative-deleted"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/initiative-deleted@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["initiative_deleted"]

[match]
intent = "Triggered when an initiative is deleted"

[[execute]]
id = "initiative"
connector = "github://ALRubinger/aileron-connector-linear"
op = "initiative_deleted"
idempotent = true
+++
