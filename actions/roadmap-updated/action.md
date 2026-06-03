+++
name = "roadmap-updated"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/roadmap-updated@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["roadmap_updated"]

[match]
intent = "Triggered when a roadmap is updated"

[[execute]]
id = "roadmap"
connector = "github://ALRubinger/aileron-connector-linear"
op = "roadmap_updated"
idempotent = true
+++
