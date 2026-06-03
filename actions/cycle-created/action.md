+++
name = "cycle-created"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/cycle-created@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["cycle_created"]

[match]
intent = "Triggered when a cycle is created"

[[execute]]
id = "cycle"
connector = "github://ALRubinger/aileron-connector-linear"
op = "cycle_created"
idempotent = true
+++
