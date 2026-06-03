+++
name = "draft-updated"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/draft-updated@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["draft_updated"]

[match]
intent = "Triggered when a draft is updated"

[[execute]]
id = "draft"
connector = "github://ALRubinger/aileron-connector-linear"
op = "draft_updated"
idempotent = true
+++
