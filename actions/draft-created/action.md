+++
name = "draft-created"
version = "0.0.0-dev"
source = "github://ALRubinger/aileron-connector-linear/actions/draft-created@0.0.0-dev"

[[requires.connectors]]
name = "github://ALRubinger/aileron-connector-linear"
version = "0.0.0-dev"
hash = "sha256:bound-at-release"
capabilities = ["draft_created"]

[match]
intent = "Triggered when a draft is created"

[[execute]]
id = "draft"
connector = "github://ALRubinger/aileron-connector-linear"
op = "draft_created"
idempotent = true
+++
